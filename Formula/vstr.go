package main

// Vstr - C string library
// Homepage: http://www.and.org/vstr/

import (
	"fmt"
	
	"os/exec"
)

func installVstr() {
	// Método 1: Descargar y extraer .tar.gz
	vstr_tar_url := "http://www.and.org/vstr/1.0.15/vstr-1.0.15.tar.bz2"
	vstr_cmd_tar := exec.Command("curl", "-L", vstr_tar_url, "-o", "package.tar.gz")
	err := vstr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vstr_zip_url := "http://www.and.org/vstr/1.0.15/vstr-1.0.15.tar.bz2"
	vstr_cmd_zip := exec.Command("curl", "-L", vstr_zip_url, "-o", "package.zip")
	err = vstr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vstr_bin_url := "http://www.and.org/vstr/1.0.15/vstr-1.0.15.tar.bz2"
	vstr_cmd_bin := exec.Command("curl", "-L", vstr_bin_url, "-o", "binary.bin")
	err = vstr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vstr_src_url := "http://www.and.org/vstr/1.0.15/vstr-1.0.15.tar.bz2"
	vstr_cmd_src := exec.Command("curl", "-L", vstr_src_url, "-o", "source.tar.gz")
	err = vstr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vstr_cmd_direct := exec.Command("./binary")
	err = vstr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
