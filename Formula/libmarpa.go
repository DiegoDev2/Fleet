package main

// Libmarpa - Marpa parse engine C library -- STABLE
// Homepage: https://jeffreykegler.github.io/Marpa-web-site/libmarpa.html

import (
	"fmt"
	
	"os/exec"
)

func installLibmarpa() {
	// Método 1: Descargar y extraer .tar.gz
	libmarpa_tar_url := "https://github.com/jeffreykegler/libmarpa/archive/refs/tags/v11.0.13.tar.gz"
	libmarpa_cmd_tar := exec.Command("curl", "-L", libmarpa_tar_url, "-o", "package.tar.gz")
	err := libmarpa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmarpa_zip_url := "https://github.com/jeffreykegler/libmarpa/archive/refs/tags/v11.0.13.zip"
	libmarpa_cmd_zip := exec.Command("curl", "-L", libmarpa_zip_url, "-o", "package.zip")
	err = libmarpa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmarpa_bin_url := "https://github.com/jeffreykegler/libmarpa/archive/refs/tags/v11.0.13.bin"
	libmarpa_cmd_bin := exec.Command("curl", "-L", libmarpa_bin_url, "-o", "binary.bin")
	err = libmarpa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmarpa_src_url := "https://github.com/jeffreykegler/libmarpa/archive/refs/tags/v11.0.13.src.tar.gz"
	libmarpa_cmd_src := exec.Command("curl", "-L", libmarpa_src_url, "-o", "source.tar.gz")
	err = libmarpa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmarpa_cmd_direct := exec.Command("./binary")
	err = libmarpa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: texlive")
	exec.Command("latte", "install", "texlive").Run()
}
