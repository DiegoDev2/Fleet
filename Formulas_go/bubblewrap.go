package main

// Bubblewrap - Unprivileged sandboxing tool for Linux
// Homepage: https://github.com/containers/bubblewrap

import (
	"fmt"
	
	"os/exec"
)

func installBubblewrap() {
	// Método 1: Descargar y extraer .tar.gz
	bubblewrap_tar_url := "https://github.com/containers/bubblewrap/releases/download/v0.10.0/bubblewrap-0.10.0.tar.xz"
	bubblewrap_cmd_tar := exec.Command("curl", "-L", bubblewrap_tar_url, "-o", "package.tar.gz")
	err := bubblewrap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bubblewrap_zip_url := "https://github.com/containers/bubblewrap/releases/download/v0.10.0/bubblewrap-0.10.0.tar.xz"
	bubblewrap_cmd_zip := exec.Command("curl", "-L", bubblewrap_zip_url, "-o", "package.zip")
	err = bubblewrap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bubblewrap_bin_url := "https://github.com/containers/bubblewrap/releases/download/v0.10.0/bubblewrap-0.10.0.tar.xz"
	bubblewrap_cmd_bin := exec.Command("curl", "-L", bubblewrap_bin_url, "-o", "binary.bin")
	err = bubblewrap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bubblewrap_src_url := "https://github.com/containers/bubblewrap/releases/download/v0.10.0/bubblewrap-0.10.0.tar.xz"
	bubblewrap_cmd_src := exec.Command("curl", "-L", bubblewrap_src_url, "-o", "source.tar.gz")
	err = bubblewrap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bubblewrap_cmd_direct := exec.Command("./binary")
	err = bubblewrap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: libxslt")
exec.Command("latte", "install", "libxslt")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: strace")
exec.Command("latte", "install", "strace")
	fmt.Println("Instalando dependencia: libcap")
exec.Command("latte", "install", "libcap")
}
