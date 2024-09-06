package main

// Kdoctor - Environment diagnostics for Kotlin Multiplatform Mobile app development
// Homepage: https://github.com/kotlin/kdoctor

import (
	"fmt"
	
	"os/exec"
)

func installKdoctor() {
	// Método 1: Descargar y extraer .tar.gz
	kdoctor_tar_url := "https://github.com/Kotlin/kdoctor/archive/refs/tags/v1.1.0.tar.gz"
	kdoctor_cmd_tar := exec.Command("curl", "-L", kdoctor_tar_url, "-o", "package.tar.gz")
	err := kdoctor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kdoctor_zip_url := "https://github.com/Kotlin/kdoctor/archive/refs/tags/v1.1.0.zip"
	kdoctor_cmd_zip := exec.Command("curl", "-L", kdoctor_zip_url, "-o", "package.zip")
	err = kdoctor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kdoctor_bin_url := "https://github.com/Kotlin/kdoctor/archive/refs/tags/v1.1.0.bin"
	kdoctor_cmd_bin := exec.Command("curl", "-L", kdoctor_bin_url, "-o", "binary.bin")
	err = kdoctor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kdoctor_src_url := "https://github.com/Kotlin/kdoctor/archive/refs/tags/v1.1.0.src.tar.gz"
	kdoctor_cmd_src := exec.Command("curl", "-L", kdoctor_src_url, "-o", "source.tar.gz")
	err = kdoctor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kdoctor_cmd_direct := exec.Command("./binary")
	err = kdoctor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
