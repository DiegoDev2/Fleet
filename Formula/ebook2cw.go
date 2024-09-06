package main

// Ebook2cw - Converts ebooks to morse code
// Homepage: https://fkurz.net/ham/ebook2cw.html

import (
	"fmt"
	
	"os/exec"
)

func installEbook2cw() {
	// Método 1: Descargar y extraer .tar.gz
	ebook2cw_tar_url := "https://fkurz.net/ham/ebook2cw/ebook2cw-0.8.5.tar.gz"
	ebook2cw_cmd_tar := exec.Command("curl", "-L", ebook2cw_tar_url, "-o", "package.tar.gz")
	err := ebook2cw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ebook2cw_zip_url := "https://fkurz.net/ham/ebook2cw/ebook2cw-0.8.5.zip"
	ebook2cw_cmd_zip := exec.Command("curl", "-L", ebook2cw_zip_url, "-o", "package.zip")
	err = ebook2cw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ebook2cw_bin_url := "https://fkurz.net/ham/ebook2cw/ebook2cw-0.8.5.bin"
	ebook2cw_cmd_bin := exec.Command("curl", "-L", ebook2cw_bin_url, "-o", "binary.bin")
	err = ebook2cw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ebook2cw_src_url := "https://fkurz.net/ham/ebook2cw/ebook2cw-0.8.5.src.tar.gz"
	ebook2cw_cmd_src := exec.Command("curl", "-L", ebook2cw_src_url, "-o", "source.tar.gz")
	err = ebook2cw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ebook2cw_cmd_direct := exec.Command("./binary")
	err = ebook2cw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
