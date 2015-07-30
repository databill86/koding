helpers = require './helpers.js'
assert  = require 'assert'

modalSelector = '.machine-settings-modal.AppModal'

module.exports =


  openVmSettingsModal: (browser, vmName, cssSelector) ->

    if not vmName
      vmName = 'koding-vm-0'

    if not cssSelector
      cssSelector = '.general'

    modalSelector = '.machine-settings-modal.AppModal'
    itemSelector  = modalSelector + ' .AppModal-navItem' + cssSelector
    vmSelector    = ".activity-sidebar .sidebar-machine-box.#{vmName} .running"

    browser
      .pause                   5000 # wait for sidebar redraw
      .waitForElementVisible   vmSelector, 20000
      .moveToElement           vmSelector + ' span', 125, 20
      .click                   vmSelector + ' span'
      .waitForElementVisible   modalSelector, 20000 # Assertion
      .pause                   2500
      .waitForElementVisible   itemSelector, 20000
      .click                   itemSelector


  openDiskUsageSettings: (browser, vmName) ->

    @openVmSettingsModal browser, vmName, '.disk-usage'


  openSpecsSettings: (browser, vmName) ->

    @openVmSettingsModal browser, vmName, '.specs'


  openGeneralSettings: (browser,vmName) ->

    @openVmSettingsModal browser, vmName


  openDomainSettings: (browser, vmName) ->

    @openVmSettingsModal browser, vmName, '.domains'

  openAdvancedSettings: (browser, vmName) ->

    @openVmSettingsModal browser, vmName, '.advanced'
  
  openSnapshotsSettings: (browser,vmName) ->

    @openVmSettingsModal browser, vmName, '.snapshots'

  clickAddVMButton: (browser) ->

    sidebarTitle = '[testpath=main-sidebar] .activity-sidebar .vms .sidebar-title'

    browser
      .waitForElementVisible   '[testpath=main-sidebar]', 20000
      .waitForElementVisible   sidebarTitle, 20000
      .moveToElement           sidebarTitle + ' a.buy-vm', 10, 10
      .click                   sidebarTitle + ' a.buy-vm'

    @clickAddKodingVMButton browser


  clickAddKodingVMButton: (browser) ->

    modalSelector  = '.environments-modal'
    buttonSelector = modalSelector + ' .kdbutton.add-vm-button'

    browser
      .waitForElementVisible   modalSelector, 20000 # Assertion
      .waitForElementVisible   buttonSelector, 20000 # Assertion
      .click                   buttonSelector


  seeUpgradeModal: (browser) ->

    @clickAddVMButton(browser)
    browser.waitForElementVisible '.computeplan-modal.free-plan .kdmodal-inner', 20000 # Assertion


  addDomain: (browser, user) ->

    buttonSelector     = '.domains .kdheaderview button.add-button'
    buttonLoaderSelector = '.add-view button.loading'
    paragraph          = helpers.getFakeText()
    createDomainName   = paragraph.split(' ')[0]
    domainName         = createDomainName + '.' + user.username + '.dev.koding.io'

    @openDomainSettings(browser)

    browser
      .waitForElementVisible    buttonSelector, 20000
      .click                    buttonSelector
      .waitForElementVisible    '.add-view input.hitenterview', 20000
      .setValue                 '.add-view input.hitenterview', createDomainName + '\n'
      .waitForElementVisible    buttonLoaderSelector, 10000
      .waitForElementNotVisible buttonLoaderSelector, 20000
      .assert.containsText      '.kdlistitemview-domain:last-child' , domainName

    return domainName


  deleteDomain: (browser, user, domainName) ->

    domainSelector = '.machine-settings-modal .kdlistitemview-domain'
    lastDomainItem = "#{domainSelector}:last-child"
    loader         = lastDomainItem + '.in-progress'

    browser
      .moveToElement             lastDomainItem, 10, 10
      .click                     lastDomainItem + ' span.remove'
      .waitForElementVisible     loader, 10000
      .waitForElementNotVisible  loader, 20000
      .getText                   domainSelector, (result) =>
        assert.notEqual          result.value, domainName # Assertion


  clickCreateVMButton: (browser) ->

    browser
      .waitForElementVisible    '.env-modal.paid-plan', 25000
      .click                    '.env-modal.paid-plan button'
      .waitForElementNotVisible '.env-modal.paid-plan', 250000
      .waitForElementVisible    '.my-machines .koding-vm-1 a:first-child', 25000
      .end()
