package main

// Fastfetch - Like neofetch, but much faster because written mostly in C
// Homepage: https://github.com/fastfetch-cli/fastfetch

import (
	"fmt"
	
	"os/exec"
)

func installFastfetch() {
	// Método 1: Descargar y extraer .tar.gz
	fastfetch_tar_url := "https://github.com/fastfetch-cli/fastfetch/archive/refs/tags/2.23.0.tar.gz"
	fastfetch_cmd_tar := exec.Command("curl", "-L", fastfetch_tar_url, "-o", "package.tar.gz")
	err := fastfetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastfetch_zip_url := "https://github.com/fastfetch-cli/fastfetch/archive/refs/tags/2.23.0.zip"
	fastfetch_cmd_zip := exec.Command("curl", "-L", fastfetch_zip_url, "-o", "package.zip")
	err = fastfetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastfetch_bin_url := "https://github.com/fastfetch-cli/fastfetch/archive/refs/tags/2.23.0.bin"
	fastfetch_cmd_bin := exec.Command("curl", "-L", fastfetch_bin_url, "-o", "binary.bin")
	err = fastfetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastfetch_src_url := "https://github.com/fastfetch-cli/fastfetch/archive/refs/tags/2.23.0.src.tar.gz"
	fastfetch_cmd_src := exec.Command("curl", "-L", fastfetch_src_url, "-o", "source.tar.gz")
	err = fastfetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastfetch_cmd_direct := exec.Command("./binary")
	err = fastfetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: chafa")
	exec.Command("latte", "install", "chafa").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vulkan-loader")
	exec.Command("latte", "install", "vulkan-loader").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: ddcutil")
	exec.Command("latte", "install", "ddcutil").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: linux-headers@5.15")
	exec.Command("latte", "install", "linux-headers@5.15").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: opencl-icd-loader")
	exec.Command("latte", "install", "opencl-icd-loader").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: rpm")
	exec.Command("latte", "install", "rpm").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
