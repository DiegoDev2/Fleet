package main

// Goredo - Go implementation of djb's redo, a Makefile replacement that sucks less
// Homepage: http://www.goredo.cypherpunks.ru/

import (
	"fmt"
	
	"os/exec"
)

func installGoredo() {
	// Método 1: Descargar y extraer .tar.gz
	goredo_tar_url := "http://www.goredo.cypherpunks.ru/download/goredo-2.6.2.tar.zst"
	goredo_cmd_tar := exec.Command("curl", "-L", goredo_tar_url, "-o", "package.tar.gz")
	err := goredo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goredo_zip_url := "http://www.goredo.cypherpunks.ru/download/goredo-2.6.2.tar.zst"
	goredo_cmd_zip := exec.Command("curl", "-L", goredo_zip_url, "-o", "package.zip")
	err = goredo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goredo_bin_url := "http://www.goredo.cypherpunks.ru/download/goredo-2.6.2.tar.zst"
	goredo_cmd_bin := exec.Command("curl", "-L", goredo_bin_url, "-o", "binary.bin")
	err = goredo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goredo_src_url := "http://www.goredo.cypherpunks.ru/download/goredo-2.6.2.tar.zst"
	goredo_cmd_src := exec.Command("curl", "-L", goredo_src_url, "-o", "source.tar.gz")
	err = goredo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goredo_cmd_direct := exec.Command("./binary")
	err = goredo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
