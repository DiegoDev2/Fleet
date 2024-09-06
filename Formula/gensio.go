package main

// Gensio - Stream I/O Library
// Homepage: https://github.com/cminyard/gensio

import (
	"fmt"
	
	"os/exec"
)

func installGensio() {
	// Método 1: Descargar y extraer .tar.gz
	gensio_tar_url := "https://github.com/cminyard/gensio/releases/download/v2.8.6/gensio-2.8.6.tar.gz"
	gensio_cmd_tar := exec.Command("curl", "-L", gensio_tar_url, "-o", "package.tar.gz")
	err := gensio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gensio_zip_url := "https://github.com/cminyard/gensio/releases/download/v2.8.6/gensio-2.8.6.zip"
	gensio_cmd_zip := exec.Command("curl", "-L", gensio_zip_url, "-o", "package.zip")
	err = gensio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gensio_bin_url := "https://github.com/cminyard/gensio/releases/download/v2.8.6/gensio-2.8.6.bin"
	gensio_cmd_bin := exec.Command("curl", "-L", gensio_bin_url, "-o", "binary.bin")
	err = gensio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gensio_src_url := "https://github.com/cminyard/gensio/releases/download/v2.8.6/gensio-2.8.6.src.tar.gz"
	gensio_cmd_src := exec.Command("curl", "-L", gensio_src_url, "-o", "source.tar.gz")
	err = gensio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gensio_cmd_direct := exec.Command("./binary")
	err = gensio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: avahi")
	exec.Command("latte", "install", "avahi").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
}
