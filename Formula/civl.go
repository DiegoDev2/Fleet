package main

// Civl - Concurrency Intermediate Verification Language
// Homepage: https://vsl.cis.udel.edu/civl/

import (
	"fmt"
	
	"os/exec"
)

func installCivl() {
	// Método 1: Descargar y extraer .tar.gz
	civl_tar_url := "https://vsl.cis.udel.edu/lib/sw/civl/1.22/r5854/release/CIVL-1.22_5854.tgz"
	civl_cmd_tar := exec.Command("curl", "-L", civl_tar_url, "-o", "package.tar.gz")
	err := civl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	civl_zip_url := "https://vsl.cis.udel.edu/lib/sw/civl/1.22/r5854/release/CIVL-1.22_5854.tgz"
	civl_cmd_zip := exec.Command("curl", "-L", civl_zip_url, "-o", "package.zip")
	err = civl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	civl_bin_url := "https://vsl.cis.udel.edu/lib/sw/civl/1.22/r5854/release/CIVL-1.22_5854.tgz"
	civl_cmd_bin := exec.Command("curl", "-L", civl_bin_url, "-o", "binary.bin")
	err = civl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	civl_src_url := "https://vsl.cis.udel.edu/lib/sw/civl/1.22/r5854/release/CIVL-1.22_5854.tgz"
	civl_cmd_src := exec.Command("curl", "-L", civl_src_url, "-o", "source.tar.gz")
	err = civl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	civl_cmd_direct := exec.Command("./binary")
	err = civl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: z3")
	exec.Command("latte", "install", "z3").Run()
}
