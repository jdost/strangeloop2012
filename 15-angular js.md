#Angular JS - HTML Redesigned for Web Apps
##Misko Hevery

Originally for web designers, to add web app-iness to their designs

Web originally designed to serve static pages, so things suck once you leave
static land

jQuery is imperitive, desired declarative

###Data Binding
Applies a model to a view, two views, image + name, fields with values of names

Get spreadsheet like updates (change A, the A+B field updates it's value)

###Advanced templating
Repeaters, many models (array) pumped into a single template, get one of the
templates evaluated against each model

Modifications of the model array should modify the same templates (swap indeces will
swap the DOM nodes)

ECMAScript is moving towards model view driven dev, Angular is "prepping" us

Angular provides fancy tags (like <tabs> and <tweet>) for an HTML DSL like thing

Reusable components is prepping us for Web Components

###Live Coding
* Uses fancy HTML5 attributes on the HTML tag
    <html ng-app> -- enables angular
* `%0` in chromium console gives highlighted element
* most of the integration with angular is through HTML attributes, sets live updates
    to the associated model
* no specific API, no DOM hitting
* built in form validation
* templates side loaded
* handles routing

AngularJS is a "framework"
"Angular is a 'better browser'"

Lookat: Testacular, distributed browser testing, using node+socketio
