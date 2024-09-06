package main

// Stanc3 - Stan transpiler
// Homepage: https://github.com/stan-dev/stanc3

import (
	"fmt"
	
	"os/exec"
)

func installStanc3() {
	// Método 1: Descargar y extraer .tar.gz
	stanc3_tar_url := "https://github.com/stan-dev/stanc3.git"
	stanc3_cmd_tar := exec.Command("curl", "-L", stanc3_tar_url, "-o", "package.tar.gz")
	err := stanc3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stanc3_zip_url := "https://github.com/stan-dev/stanc3.git"
	stanc3_cmd_zip := exec.Command("curl", "-L", stanc3_zip_url, "-o", "package.zip")
	err = stanc3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stanc3_bin_url := "https://github.com/stan-dev/stanc3.git"
	stanc3_cmd_bin := exec.Command("curl", "-L", stanc3_bin_url, "-o", "binary.bin")
	err = stanc3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stanc3_src_url := "https://github.com/stan-dev/stanc3.git"
	stanc3_cmd_src := exec.Command("curl", "-L", stanc3_src_url, "-o", "source.tar.gz")
	err = stanc3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stanc3_cmd_direct := exec.Command("./binary")
	err = stanc3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
}
