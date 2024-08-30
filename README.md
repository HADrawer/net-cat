# net-cat

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/almadhoob/net-cat">
    <img src="logo.png" alt="Logo" width="256" height="256">
  </a>

<h3 align="center">NetCat Server-Client</h3>

  <p align="center">
    A simple Unix terminal app for internal communication.
    <br />
    <a href="https://github.com/almadhoob/net-cat"><strong>Go to the repo »</strong></a>
    <br />
    <br />
    <a href="https://github.com/almadhoob/net-cat/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/almadhoob/net-cat/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#authors">Authors</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

<!-- <div align="center"><img src="images/screenshot.png" alt="Screenshot"></div> -->
<!-- <br /> -->

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server. It is written in Golang using the following standard packages: io, log, os, fmt, net, sync, time, bufio, errors, strings, and reflect.

<p align="right">(<a href="#net-cat">back to top</a>)</p>

### Built With

- [Go programming language](https://go.dev/doc/)
- [GOCUI console user interface](https://pkg.go.dev/github.com/jroimartin/gocui/)

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

You only need a web browser (such as Mozilla Firefox) besides the following software:

- Go programming language (v1.23.0 or newer)
  ```sh
  go version
  ```

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/almadhoob/net-cat.git
   cd net-cat
   ```

2. Build the app

   ```sh
   go build -o app.run
   chmod +x app.run
   ```

3. Run the app

   ```sh
   ./app.run
   ```

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

You can run the app as the following:

   ```sh
   go run . 2525 localhost
   ```

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Added some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

Don't forget to give the project a star! Thanks again!

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- AUTHORS -->

## Authors

- Hashem Ahmed — [GitHub](https://github.com/hadrawer)
- Ahmed Almadhoob — [GitHub](https://github.com/almadhoob)

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- Yaman Al-Masri
- Ahmed Abdeen

<p align="right">(<a href="#net-cat">back to top</a>)</p>

<!-- LICENSE -->

## License

This is an [MIT-Licensed](./LICENSE) project which is created by its authors for [Reboot01](https://reboot01.com/).

<p align="right">(<a href="#net-cat">back to top</a>)</p>
