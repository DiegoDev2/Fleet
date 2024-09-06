package main

// Cabocha - Yet Another Japanese Dependency Structure Analyzer
// Homepage: https://taku910.github.io/cabocha/

import (
	"fmt"
	
	"os/exec"
)

func installCabocha() {
	// Método 1: Descargar y extraer .tar.gz
	cabocha_tar_url := "https://distfiles.macports.org/cabocha/cabocha-0.69.tar.bz2"
	cabocha_cmd_tar := exec.Command("curl", "-L", cabocha_tar_url, "-o", "package.tar.gz")
	err := cabocha_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cabocha_zip_url := "https://distfiles.macports.org/cabocha/cabocha-0.69.tar.bz2"
	cabocha_cmd_zip := exec.Command("curl", "-L", cabocha_zip_url, "-o", "package.zip")
	err = cabocha_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cabocha_bin_url := "https://distfiles.macports.org/cabocha/cabocha-0.69.tar.bz2"
	cabocha_cmd_bin := exec.Command("curl", "-L", cabocha_bin_url, "-o", "binary.bin")
	err = cabocha_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cabocha_src_url := "https://distfiles.macports.org/cabocha/cabocha-0.69.tar.bz2"
	cabocha_cmd_src := exec.Command("curl", "-L", cabocha_src_url, "-o", "source.tar.gz")
	err = cabocha_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cabocha_cmd_direct := exec.Command("./binary")
	err = cabocha_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: crf++")
	exec.Command("latte", "install", "crf++").Run()
	fmt.Println("Instalando dependencia: mecab")
	exec.Command("latte", "install", "mecab").Run()
	fmt.Println("Instalando dependencia: mecab-ipadic")
	exec.Command("latte", "install", "mecab-ipadic").Run()
}
