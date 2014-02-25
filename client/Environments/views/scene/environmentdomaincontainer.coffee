class EnvironmentDomainContainer extends EnvironmentContainer

  EnvironmentDataProvider.addProvider "domains", ->

    new Promise (resolve, reject)->

      KD.whoami().fetchDomains (err, domains)->

        if err or not domains or domains.length is 0
          warn "Failed to fetch domains", err  if err
          return resolve []

        resolve domains

  constructor:(options={}, data)->
    options.cssClass  = 'domains'
    options.itemClass = EnvironmentDomainItem
    options.title     = 'domains'
    super options, data

    # Plus button on domainsContainer opens up the domainCreateModal
    @on 'PlusButtonClicked', =>
      return unless KD.isLoggedIn()
        new KDNotificationView title: "You need to login to add a new domain."

      domainCreateForm = @getDomainCreateForm()

      new KDModalView
        title          : "Add Domain"
        view           : domainCreateForm
        width          : 700
        buttons        :
          createButton :
            title      : "Create"
            style      : "modal-clean-green"
            callback   : =>
              domainCreateForm.createSubDomain()

  addDomain: (domain)->

    @addItem
      title       : domain.domain
      description : $.timeago domain.createdAt
      activated   : yes
      aliases     : domain.hostnameAlias
      domain      : domain

  getDomainCreateForm: ->
    domainCreateForm = new DomainCreateForm

    @on "itemRemoved", domainCreateForm.bound "updateDomains"
    domainCreateForm.on "DomainSaved", (domain) =>
      @addDomain domain
      @emit "itemAdded"

    return domainCreateForm
