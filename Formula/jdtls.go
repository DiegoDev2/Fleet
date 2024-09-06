package main

// Jdtls - Java language specific implementation of the Language Server Protocol
// Homepage: https://github.com/eclipse-jdtls/eclipse.jdt.ls

import (
	"fmt"
	
	"os/exec"
)

func installJdtls() {
	// Método 1: Descargar y extraer .tar.gz
	jdtls_tar_url := "https://www.eclipse.org/downloads/download.php?file=/jdtls/milestones/1.39.0/jdt-language-server-1.39.0-202408291433.tar.gz"
	jdtls_cmd_tar := exec.Command("curl", "-L", jdtls_tar_url, "-o", "package.tar.gz")
	err := jdtls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jdtls_zip_url := "https://www.eclipse.org/downloads/download.php?file=/jdtls/milestones/1.39.0/jdt-language-server-1.39.0-202408291433.zip"
	jdtls_cmd_zip := exec.Command("curl", "-L", jdtls_zip_url, "-o", "package.zip")
	err = jdtls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jdtls_bin_url := "https://www.eclipse.org/downloads/download.php?file=/jdtls/milestones/1.39.0/jdt-language-server-1.39.0-202408291433.bin"
	jdtls_cmd_bin := exec.Command("curl", "-L", jdtls_bin_url, "-o", "binary.bin")
	err = jdtls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jdtls_src_url := "https://www.eclipse.org/downloads/download.php?file=/jdtls/milestones/1.39.0/jdt-language-server-1.39.0-202408291433.src.tar.gz"
	jdtls_cmd_src := exec.Command("curl", "-L", jdtls_src_url, "-o", "source.tar.gz")
	err = jdtls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jdtls_cmd_direct := exec.Command("./binary")
	err = jdtls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
