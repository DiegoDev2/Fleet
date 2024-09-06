package main

// Voldemort - Distributed key-value storage system
// Homepage: https://www.project-voldemort.com/

import (
	"fmt"
	
	"os/exec"
)

func installVoldemort() {
	// Método 1: Descargar y extraer .tar.gz
	voldemort_tar_url := "https://github.com/voldemort/voldemort/archive/refs/tags/release-1.10.26-cutoff.tar.gz"
	voldemort_cmd_tar := exec.Command("curl", "-L", voldemort_tar_url, "-o", "package.tar.gz")
	err := voldemort_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	voldemort_zip_url := "https://github.com/voldemort/voldemort/archive/refs/tags/release-1.10.26-cutoff.zip"
	voldemort_cmd_zip := exec.Command("curl", "-L", voldemort_zip_url, "-o", "package.zip")
	err = voldemort_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	voldemort_bin_url := "https://github.com/voldemort/voldemort/archive/refs/tags/release-1.10.26-cutoff.bin"
	voldemort_cmd_bin := exec.Command("curl", "-L", voldemort_bin_url, "-o", "binary.bin")
	err = voldemort_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	voldemort_src_url := "https://github.com/voldemort/voldemort/archive/refs/tags/release-1.10.26-cutoff.src.tar.gz"
	voldemort_cmd_src := exec.Command("curl", "-L", voldemort_src_url, "-o", "source.tar.gz")
	err = voldemort_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	voldemort_cmd_direct := exec.Command("./binary")
	err = voldemort_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
