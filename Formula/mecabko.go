package main

// MecabKo - See mecab
// Homepage: https://bitbucket.org/eunjeon/mecab-ko

import (
	"fmt"
	
	"os/exec"
)

func installMecabKo() {
	// Método 1: Descargar y extraer .tar.gz
	mecabko_tar_url := "https://bitbucket.org/eunjeon/mecab-ko/downloads/mecab-0.996-ko-0.9.2.tar.gz"
	mecabko_cmd_tar := exec.Command("curl", "-L", mecabko_tar_url, "-o", "package.tar.gz")
	err := mecabko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabko_zip_url := "https://bitbucket.org/eunjeon/mecab-ko/downloads/mecab-0.996-ko-0.9.2.zip"
	mecabko_cmd_zip := exec.Command("curl", "-L", mecabko_zip_url, "-o", "package.zip")
	err = mecabko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabko_bin_url := "https://bitbucket.org/eunjeon/mecab-ko/downloads/mecab-0.996-ko-0.9.2.bin"
	mecabko_cmd_bin := exec.Command("curl", "-L", mecabko_bin_url, "-o", "binary.bin")
	err = mecabko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabko_src_url := "https://bitbucket.org/eunjeon/mecab-ko/downloads/mecab-0.996-ko-0.9.2.src.tar.gz"
	mecabko_cmd_src := exec.Command("curl", "-L", mecabko_src_url, "-o", "source.tar.gz")
	err = mecabko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabko_cmd_direct := exec.Command("./binary")
	err = mecabko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
