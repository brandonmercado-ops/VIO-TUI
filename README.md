<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![License][license-shield]][license-url]

<br />
<div align="center">

  <div align="center">
    <img src="docs/assets/images/vio_github_logo.webp" alt="VIO-TUI Logo" width="400">
  </div>

  <p align="center">
    Terminal User Interface for daily time-management tracking — built by a student software engineer, for student software engineers.
    <br />
    <a href="https://github.com/whitegunrose/VIO-TUI/wiki/Documentation" target="_blank"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/whitegunrose/VIO-TUI/issues/new?labels=bug&template=bug-report.md" target="_blank">Report Bug</a>
    ·
    <a href="https://github.com/whitegunrose/VIO-TUI/issues/new?labels=enhancement&template=feature-request.md" target="_blank">Request Feature</a>
  </p>
</div>

---

## Table of Contents

1. [About The Project](#about-the-project)
   - [Built With](#built-with)
2. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
3. [Usage](#usage)
4. [Roadmap](#roadmap)
5. [Contributing](#contributing)
6. [License](#license)
7. [Contact](#contact)
8. [Acknowledgments](#acknowledgments)

---
![](https://github.com/whitegunrose/VIO-TUI/blob/main/docs/assets/demos/demo.gif)
---

## About The Project

<!-- Replace the line below with an actual screenshot once available -->
<!-- ![VIO-TUI Screenshot](docs/screenshot.png) -->

VIO-TUI is a keyboard-driven Terminal User Interface application written in Go, designed to help student software engineers take control of their daily schedules. Managing lectures, coding sessions, deadlines, and personal tasks across multiple apps is overwhelming — VIO solves that by keeping everything in one distraction-free terminal window.

**Why VIO?**

- 🖥️ Stays out of your way — no GUI overhead, no mouse required
- ⚡ Instant startup and near-zero resource usage, so it never slows down your workflow
- 📋 Tracks your time blocks, tasks, and sessions in one unified view
- 🎓 Tailored to the rhythms of a student engineering life — classes, side projects, and deadlines all in one place

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Go][Go-badge]][Go-url]
[![tview][tview-badge]][tview-url]

> VIO-TUI is 100% Go. The TUI layer is powered by [tview](https://github.com/rivo/tview) — a rich interactive widget library built on top of [tcell](https://github.com/gdamore/tcell).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Getting Started

Follow the steps below to get VIO-TUI running on your local machine.

### Prerequisites

- **Go 1.21+** — [Install Go](https://go.dev/dl/)

  ```sh
  go version   # verify installation
  ```

### Installation

1. TODO

<!-- 1. Clone the repository -->
<!---->
<!--    ```sh -->
<!--    git clone https://github.com/whitegunrose/VIO-TUI.git -->
<!--    ``` -->
<!---->
<!-- 2. Navigate into the project directory -->
<!---->
<!--    ```sh -->
<!--    cd VIO-TUI -->
<!--    ``` -->
<!---->
<!-- 3. Build the application -->
<!---->
<!--    ```sh -->
<!--    go build -o vio ./vio-tui -->
<!--    ``` -->
<!---->
<!-- 4. Run VIO -->
<!---->
<!--    ```sh -->
<!--    ./vio -->
<!--    ``` -->
<!---->
<!--    Or run it directly without building: -->
<!---->
<!--    ```sh -->
<!--    go run ./vio-tui -->
<!--    ``` -->
<!---->
<!-- 5. *(Optional)* Move the binary to your `PATH` for system-wide access -->
<!---->
<!--    ```sh -->
<!--    mv vio /usr/local/bin/vio -->
<!--    ``` -->
<!---->
<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Usage

Launch VIO from any terminal and use the keyboard to navigate:

1. TODO
<!---->
<!-- ``` -->
<!-- vio -->
<!-- ``` -->
<!---->
<!-- | Key          | Action                        | -->
<!-- |--------------|-------------------------------| -->
<!-- | `↑` / `↓`   | Navigate between entries      | -->
<!-- | `n`          | Create a new time block/task  | -->
<!-- | `e`          | Edit the selected entry       | -->
<!-- | `d`          | Delete the selected entry     | -->
<!-- | `q` / `Ctrl+C` | Quit VIO                   | -->
<!---->
<!-- > For full keybinding reference and configuration options, see the [documentation](https://github.com/whitegunrose/VIO-TUI/tree/main/docs). -->
<!---->
<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Roadmap

- [x] Core TUI shell and navigation
- [x] Daily task/time-block tracking
- [x] Canvas integration to import classes and assignments
- [ ] Persistent storage (save sessions across restarts)
- [ ] Weekly and monthly summary views
- [ ] Cross-platform binary releases (Linux, macOS, Windows)

<!-- See the [open issues](https://github.com/whitegunrose/VIO-TUI/issues) for the full list of proposed features and known bugs. -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contributing

1. TODO
<!-- Contributions are what make the open-source community such an amazing place to learn, grow, and build. Any contributions you make are **greatly appreciated**. -->
<!---->
<!-- If you have a suggestion that would improve VIO, please fork the repo and open a pull request, or open an issue with the tag `enhancement`. Don't forget to give the project a ⭐ — it means a lot! -->
<!---->
<!-- 1. Fork the project -->
<!-- 2. Create your feature branch -->
<!---->
<!--    ```sh -->
<!--    git checkout -b feature/AmazingFeature -->
<!--    ``` -->
<!---->
<!-- 3. Commit your changes -->
<!---->
<!--    ```sh -->
<!--    git commit -m 'Add some AmazingFeature' -->
<!--    ``` -->
<!---->
<!-- 4. Push to the branch -->
<!---->
<!--    ```sh -->
<!--    git push origin feature/AmazingFeature -->
<!--    ``` -->
<!---->
<!-- 5. Open a Pull Request -->
<!---->
<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Contact

**whitegunrose** — [@whitegunrose](https://github.com/whitegunrose)

<!-- Project Link: [https://github.com/whitegunrose/VIO-TUI](https://github.com/whitegunrose/VIO-TUI) -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

## Acknowledgments

Resources and tools that made VIO-TUI possible:

- [tview](https://github.com/rivo/tview) — Rich interactive TUI widget library for Go
- [tcell](https://github.com/gdamore/tcell) — Low-level terminal handling that powers tview
- [Go Documentation](https://go.dev/doc/)
- [Best-README-Template](https://github.com/othneildrew/Best-README-Template) — README structure inspiration
- [Shields.io](https://shields.io) — Badge generation
- [Choose an Open Source License](https://choosealicense.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- MARKDOWN LINKS & BADGES -->
[contributors-shield]: https://img.shields.io/github/contributors/whitegunrose/VIO-TUI.svg?style=for-the-badge
[contributors-url]: https://github.com/whitegunrose/VIO-TUI/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/whitegunrose/VIO-TUI.svg?style=for-the-badge
[forks-url]: https://github.com/whitegunrose/VIO-TUI/network/members
[stars-shield]: https://img.shields.io/github/stars/whitegunrose/VIO-TUI.svg?style=for-the-badge
[stars-url]: https://github.com/whitegunrose/VIO-TUI/stargazers
[issues-shield]: https://img.shields.io/github/issues/whitegunrose/VIO-TUI.svg?style=for-the-badge
[issues-url]: https://github.com/whitegunrose/VIO-TUI/issues
[license-shield]: https://img.shields.io/github/license/whitegunrose/VIO-TUI.svg?style=for-the-badge
[license-url]: https://github.com/whitegunrose/VIO-TUI/blob/main/LICENSE
[Go-badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[tview-badge]: https://img.shields.io/badge/tview-00ADD8?style=for-the-badge&logo=go&logoColor=white
[tview-url]: https://github.com/rivo/tview
