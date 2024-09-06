package main

// MecabIpadic - IPA dictionary compiled for MeCab
// Homepage: https://taku910.github.io/mecab/

import (
	"fmt"
	
	"os/exec"
)

func installMecabIpadic() {
	// Método 1: Descargar y extraer .tar.gz
	mecabipadic_tar_url := "https://deb.debian.org/debian/pool/main/m/mecab-ipadic/mecab-ipadic_2.7.0-20070801+main.orig.tar.gz"
	mecabipadic_cmd_tar := exec.Command("curl", "-L", mecabipadic_tar_url, "-o", "package.tar.gz")
	err := mecabipadic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabipadic_zip_url := "https://deb.debian.org/debian/pool/main/m/mecab-ipadic/mecab-ipadic_2.7.0-20070801+main.orig.zip"
	mecabipadic_cmd_zip := exec.Command("curl", "-L", mecabipadic_zip_url, "-o", "package.zip")
	err = mecabipadic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabipadic_bin_url := "https://deb.debian.org/debian/pool/main/m/mecab-ipadic/mecab-ipadic_2.7.0-20070801+main.orig.bin"
	mecabipadic_cmd_bin := exec.Command("curl", "-L", mecabipadic_bin_url, "-o", "binary.bin")
	err = mecabipadic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabipadic_src_url := "https://deb.debian.org/debian/pool/main/m/mecab-ipadic/mecab-ipadic_2.7.0-20070801+main.orig.src.tar.gz"
	mecabipadic_cmd_src := exec.Command("curl", "-L", mecabipadic_src_url, "-o", "source.tar.gz")
	err = mecabipadic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabipadic_cmd_direct := exec.Command("./binary")
	err = mecabipadic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mecab")
	exec.Command("latte", "install", "mecab").Run()
}
