LoginInputView     = require './../login/logininputview'
RegisterInlineForm = require './../login/registerform'


module.exports = class HomeRegisterForm extends RegisterInlineForm

  ENTER          = 13

  constructor:(options={},data)->

    super options, data

    @email.destroy()
    @email = new LoginInputView
      inputOptions    :
        name          : "email"
        placeholder   : "Email address"
        testPath      : "register-form-email"
        keyup         : (event) => @button.click event  if event.which is ENTER
        validate      : @getEmailValidator()
        decorateValidation: no
        focus         : => @email.icon.unsetTooltip()
        keyup         : (event) => @submitForm event  if event.which is ENTER

    @password.destroy()
    @password = new LoginInputView
      inputOptions    :
        name          : "password"
        type          : "password"
        testPath      : "recover-password"
        placeholder   : "Password"
        keyup         : (event) =>
          if event.which is ENTER
            @password.input.validate()
            @button.click event
        validate      :
          container   : this
          rules       :
            required  : yes
            minLength : 8
          messages    :
            required  : 'Please enter a password.'
            minLength : 'Passwords should be at least 8 characters.'
          events      :
            required  : 'blur'
            minLength : 'blur'

    @email.setOption 'stickyTooltip', yes
    @password.setOption 'stickyTooltip', yes

    @email.input.on    'focus', @bound 'handleFocus'
    @password.input.on 'focus', @bound 'handleFocus'
    @email.input.on    'blur',  => @fetchGravatarInfo @email.input.getValue()

    KD.singletons.router.on 'RouteInfoHandled', =>
      @email.icon.unsetTooltip()
      @password.icon.unsetTooltip()

    @addCustomData 'redirectTo', 'Hackathon/Apply'


  handleFocus: ->

    super

    video = document.getElementById 'bgVideo'

    speed = 1

    repeater = KD.utils.repeat 20, ->

      speed -= .02
      video.playbackRate = speed

      if speed <= 0
        video.pause()
        KD.utils.killRepeat repeater


  pistachio : ->

    """
    <section class='clearfix'>
      <div class='fl email'>{{> @email}}</div>
      <div class='fl password'>{{> @password}}</div>
      <div class='fl submit'>{{> @button}}</div>
    </section>
    """
