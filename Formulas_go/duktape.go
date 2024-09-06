package main

// Duktape - Embeddable Javascript engine with compact footprint
// Homepage: https://duktape.org

import (
	"fmt"
	
	"os/exec"
)

func installDuktape() {
	// Método 1: Descargar y extraer .tar.gz
	duktape_tar_url := "https://github.com/svaarala/duktape/releases/download/v2.7.0/duktape-2.7.0.tar.xz"
	duktape_cmd_tar := exec.Command("curl", "-L", duktape_tar_url, "-o", "package.tar.gz")
	err := duktape_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duktape_zip_url := "https://github.com/svaarala/duktape/releases/download/v2.7.0/duktape-2.7.0.tar.xz"
	duktape_cmd_zip := exec.Command("curl", "-L", duktape_zip_url, "-o", "package.zip")
	err = duktape_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duktape_bin_url := "https://github.com/svaarala/duktape/releases/download/v2.7.0/duktape-2.7.0.tar.xz"
	duktape_cmd_bin := exec.Command("curl", "-L", duktape_bin_url, "-o", "binary.bin")
	err = duktape_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duktape_src_url := "https://github.com/svaarala/duktape/releases/download/v2.7.0/duktape-2.7.0.tar.xz"
	duktape_cmd_src := exec.Command("curl", "-L", duktape_src_url, "-o", "source.tar.gz")
	err = duktape_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duktape_cmd_direct := exec.Command("./binary")
	err = duktape_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
