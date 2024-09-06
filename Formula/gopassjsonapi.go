package main

// GopassJsonapi - Gopass Browser Bindings
// Homepage: https://github.com/gopasspw/gopass-jsonapi

import (
	"fmt"
	
	"os/exec"
)

func installGopassJsonapi() {
	// Método 1: Descargar y extraer .tar.gz
	gopassjsonapi_tar_url := "https://github.com/gopasspw/gopass-jsonapi/archive/refs/tags/v1.15.14.tar.gz"
	gopassjsonapi_cmd_tar := exec.Command("curl", "-L", gopassjsonapi_tar_url, "-o", "package.tar.gz")
	err := gopassjsonapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gopassjsonapi_zip_url := "https://github.com/gopasspw/gopass-jsonapi/archive/refs/tags/v1.15.14.zip"
	gopassjsonapi_cmd_zip := exec.Command("curl", "-L", gopassjsonapi_zip_url, "-o", "package.zip")
	err = gopassjsonapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gopassjsonapi_bin_url := "https://github.com/gopasspw/gopass-jsonapi/archive/refs/tags/v1.15.14.bin"
	gopassjsonapi_cmd_bin := exec.Command("curl", "-L", gopassjsonapi_bin_url, "-o", "binary.bin")
	err = gopassjsonapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gopassjsonapi_src_url := "https://github.com/gopasspw/gopass-jsonapi/archive/refs/tags/v1.15.14.src.tar.gz"
	gopassjsonapi_cmd_src := exec.Command("curl", "-L", gopassjsonapi_src_url, "-o", "source.tar.gz")
	err = gopassjsonapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gopassjsonapi_cmd_direct := exec.Command("./binary")
	err = gopassjsonapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: gopass")
	exec.Command("latte", "install", "gopass").Run()
}
