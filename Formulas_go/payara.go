package main

// Payara - Java EE application server forked from GlassFish
// Homepage: https://www.payara.fish

import (
	"fmt"
	
	"os/exec"
)

func installPayara() {
	// Método 1: Descargar y extraer .tar.gz
	payara_tar_url := "https://search.maven.org/remotecontent?filepath=fish/payara/distributions/payara/6.2024.8/payara-6.2024.8.zip"
	payara_cmd_tar := exec.Command("curl", "-L", payara_tar_url, "-o", "package.tar.gz")
	err := payara_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	payara_zip_url := "https://search.maven.org/remotecontent?filepath=fish/payara/distributions/payara/6.2024.8/payara-6.2024.8.zip"
	payara_cmd_zip := exec.Command("curl", "-L", payara_zip_url, "-o", "package.zip")
	err = payara_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	payara_bin_url := "https://search.maven.org/remotecontent?filepath=fish/payara/distributions/payara/6.2024.8/payara-6.2024.8.zip"
	payara_cmd_bin := exec.Command("curl", "-L", payara_bin_url, "-o", "binary.bin")
	err = payara_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	payara_src_url := "https://search.maven.org/remotecontent?filepath=fish/payara/distributions/payara/6.2024.8/payara-6.2024.8.zip"
	payara_cmd_src := exec.Command("curl", "-L", payara_src_url, "-o", "source.tar.gz")
	err = payara_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	payara_cmd_direct := exec.Command("./binary")
	err = payara_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
