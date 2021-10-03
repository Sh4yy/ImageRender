![thumbnail](https://user-images.githubusercontent.com/23535123/135739086-3d21ce14-6f00-42d3-8b6b-b56c625be83f.jpg)
## Render dynamic images from html template on demand

Create a template inside the ```/templates``` directory and add your static content to ```/static```
You can populate and generate images via making a get request to the ```/:template``` route and passing your template variables as query parameter.

#### Example
The template for the image above already exist in the repo as ```thumbnail.html```. You can render making a get request to ```http://localhost:8080/thumbnail```
and passing the following url-encoded query parameters:

```title: Render dynamic images from html templates```\
```user: Shayan```\
```quality: 100``` optional\
```width: 1200``` optional\
```height: 627``` optional

Example:\
```http://localhost:8080/thumbnail?title=Render%20dynamic%20images%20from%20html%20templates&user=Shayan%20Taslim&quality=100&width=1200&height=672&format=jpeg```
