package main

// Alot - Text mode MUA using notmuch mail
// Homepage: https://github.com/pazz/alot

import (
	"fmt"
	
	"os/exec"
)

func installAlot() {
	// Método 1: Descargar y extraer .tar.gz
	alot_tar_url := "https://github.com/pazz/alot/archive/refs/tags/0.10.tar.gz"
	alot_cmd_tar := exec.Command("curl", "-L", alot_tar_url, "-o", "package.tar.gz")
	err := alot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alot_zip_url := "https://github.com/pazz/alot/archive/refs/tags/0.10.zip"
	alot_cmd_zip := exec.Command("curl", "-L", alot_zip_url, "-o", "package.zip")
	err = alot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alot_bin_url := "https://github.com/pazz/alot/archive/refs/tags/0.10.bin"
	alot_cmd_bin := exec.Command("curl", "-L", alot_bin_url, "-o", "binary.bin")
	err = alot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alot_src_url := "https://github.com/pazz/alot/archive/refs/tags/0.10.src.tar.gz"
	alot_cmd_src := exec.Command("curl", "-L", alot_src_url, "-o", "source.tar.gz")
	err = alot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alot_cmd_direct := exec.Command("./binary")
	err = alot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: notmuch")
	exec.Command("latte", "install", "notmuch").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
