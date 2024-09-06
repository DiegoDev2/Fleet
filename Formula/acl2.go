package main

// Acl2 - Logic and programming language in which you can model computer systems
// Homepage: https://www.cs.utexas.edu/users/moore/acl2/index.html

import (
	"fmt"
	
	"os/exec"
)

func installAcl2() {
	// Método 1: Descargar y extraer .tar.gz
	acl2_tar_url := "https://github.com/acl2/acl2/archive/refs/tags/8.5.tar.gz"
	acl2_cmd_tar := exec.Command("curl", "-L", acl2_tar_url, "-o", "package.tar.gz")
	err := acl2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	acl2_zip_url := "https://github.com/acl2/acl2/archive/refs/tags/8.5.zip"
	acl2_cmd_zip := exec.Command("curl", "-L", acl2_zip_url, "-o", "package.zip")
	err = acl2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	acl2_bin_url := "https://github.com/acl2/acl2/archive/refs/tags/8.5.bin"
	acl2_cmd_bin := exec.Command("curl", "-L", acl2_bin_url, "-o", "binary.bin")
	err = acl2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	acl2_src_url := "https://github.com/acl2/acl2/archive/refs/tags/8.5.src.tar.gz"
	acl2_cmd_src := exec.Command("curl", "-L", acl2_src_url, "-o", "source.tar.gz")
	err = acl2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	acl2_cmd_direct := exec.Command("./binary")
	err = acl2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbcl")
	exec.Command("latte", "install", "sbcl").Run()
}
