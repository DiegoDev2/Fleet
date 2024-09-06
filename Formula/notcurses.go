package main

// Notcurses - Blingful character graphics/TUI library
// Homepage: https://nick-black.com/dankwiki/index.php/Notcurses

import (
	"fmt"
	
	"os/exec"
)

func installNotcurses() {
	// Método 1: Descargar y extraer .tar.gz
	notcurses_tar_url := "https://github.com/dankamongmen/notcurses/archive/refs/tags/v3.0.9.tar.gz"
	notcurses_cmd_tar := exec.Command("curl", "-L", notcurses_tar_url, "-o", "package.tar.gz")
	err := notcurses_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	notcurses_zip_url := "https://github.com/dankamongmen/notcurses/archive/refs/tags/v3.0.9.zip"
	notcurses_cmd_zip := exec.Command("curl", "-L", notcurses_zip_url, "-o", "package.zip")
	err = notcurses_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	notcurses_bin_url := "https://github.com/dankamongmen/notcurses/archive/refs/tags/v3.0.9.bin"
	notcurses_cmd_bin := exec.Command("curl", "-L", notcurses_bin_url, "-o", "binary.bin")
	err = notcurses_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	notcurses_src_url := "https://github.com/dankamongmen/notcurses/archive/refs/tags/v3.0.9.src.tar.gz"
	notcurses_cmd_src := exec.Command("curl", "-L", notcurses_src_url, "-o", "source.tar.gz")
	err = notcurses_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	notcurses_cmd_direct := exec.Command("./binary")
	err = notcurses_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doctest")
	exec.Command("latte", "install", "doctest").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: libdeflate")
	exec.Command("latte", "install", "libdeflate").Run()
	fmt.Println("Instalando dependencia: libunistring")
	exec.Command("latte", "install", "libunistring").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
}
