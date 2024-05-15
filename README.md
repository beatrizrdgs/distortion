<a name="readme-top"></a>


<div align="center">

[![Stargazers][stars-shield]][stars-url] [![Issues][issues-shield]][issues-url] [![MIT License][license-shield]][license-url]

</div>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/beatrizrdgs/distortion">
    <img src="https://i.imgur.com/BiySi5V.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Distortion</h3>

  <p align="center">
    App for image manipulation
    <br />
    <br />
    <a href="">View Demo</a>
    ·
    <a href="https://github.com/beatrizrdgs/distortion/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/beatrizrdgs/distortion/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![Distortion Screen Shot][product-screenshot]

Distortion is a web-based tool designed for manipulating images. It offers a user-friendly interface to apply various transformations to images.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- Usage -->
## Usage

### Running locally

You can run the project with either Go or Docker. Please note that this assumes you already have Go or Docker installed.

**Go**
```bash
go run main.go server
```


**Docker**
```bash
docker build -t distortion-img .
docker run --network host -d --name distortion-cont -p 8080:8080 distortion-img
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GNU Affero General Public License v3.0. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[stars-shield]: https://img.shields.io/github/stars/beatrizrdgs/distortion.svg?style=for-the-badge
[stars-url]: https://github.com/beatrizrdgs/distortion/stargazers
[issues-shield]: https://img.shields.io/github/issues/beatrizrdgs/distortion.svg?style=for-the-badge
[issues-url]: https://github.com/beatrizrdgs/distortion/issues
[license-shield]: https://img.shields.io/github/license/beatrizrdgs/distortion.svg?style=for-the-badge
[license-url]: https://github.com/beatrizrdgs/distortion/blob/master/LICENSE
[product-screenshot]: https://i.imgur.com/SQa6lPi.png