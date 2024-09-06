package main

// Mecab - Yet another part-of-speech and morphological analyzer
// Homepage: https://taku910.github.io/mecab/

import (
	"fmt"
	
	"os/exec"
)

func installMecab() {
	// Método 1: Descargar y extraer .tar.gz
	mecab_tar_url := "https://deb.debian.org/debian/pool/main/m/mecab/mecab_0.996.orig.tar.gz"
	mecab_cmd_tar := exec.Command("curl", "-L", mecab_tar_url, "-o", "package.tar.gz")
	err := mecab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecab_zip_url := "https://deb.debian.org/debian/pool/main/m/mecab/mecab_0.996.orig.zip"
	mecab_cmd_zip := exec.Command("curl", "-L", mecab_zip_url, "-o", "package.zip")
	err = mecab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecab_bin_url := "https://deb.debian.org/debian/pool/main/m/mecab/mecab_0.996.orig.bin"
	mecab_cmd_bin := exec.Command("curl", "-L", mecab_bin_url, "-o", "binary.bin")
	err = mecab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecab_src_url := "https://deb.debian.org/debian/pool/main/m/mecab/mecab_0.996.orig.src.tar.gz"
	mecab_cmd_src := exec.Command("curl", "-L", mecab_src_url, "-o", "source.tar.gz")
	err = mecab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecab_cmd_direct := exec.Command("./binary")
	err = mecab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
