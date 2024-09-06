package main

// Kepubify - Convert ebooks from epub to kepub
// Homepage: https://pgaskin.net/kepubify/

import (
	"fmt"
	
	"os/exec"
)

func installKepubify() {
	// Método 1: Descargar y extraer .tar.gz
	kepubify_tar_url := "https://github.com/pgaskin/kepubify/archive/refs/tags/v4.0.4.tar.gz"
	kepubify_cmd_tar := exec.Command("curl", "-L", kepubify_tar_url, "-o", "package.tar.gz")
	err := kepubify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kepubify_zip_url := "https://github.com/pgaskin/kepubify/archive/refs/tags/v4.0.4.zip"
	kepubify_cmd_zip := exec.Command("curl", "-L", kepubify_zip_url, "-o", "package.zip")
	err = kepubify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kepubify_bin_url := "https://github.com/pgaskin/kepubify/archive/refs/tags/v4.0.4.bin"
	kepubify_cmd_bin := exec.Command("curl", "-L", kepubify_bin_url, "-o", "binary.bin")
	err = kepubify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kepubify_src_url := "https://github.com/pgaskin/kepubify/archive/refs/tags/v4.0.4.src.tar.gz"
	kepubify_cmd_src := exec.Command("curl", "-L", kepubify_src_url, "-o", "source.tar.gz")
	err = kepubify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kepubify_cmd_direct := exec.Command("./binary")
	err = kepubify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
