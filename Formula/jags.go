package main

// Jags - Just Another Gibbs Sampler for Bayesian MCMC simulation
// Homepage: https://mcmc-jags.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installJags() {
	// Método 1: Descargar y extraer .tar.gz
	jags_tar_url := "https://downloads.sourceforge.net/project/mcmc-jags/JAGS/4.x/Source/JAGS-4.3.2.tar.gz"
	jags_cmd_tar := exec.Command("curl", "-L", jags_tar_url, "-o", "package.tar.gz")
	err := jags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jags_zip_url := "https://downloads.sourceforge.net/project/mcmc-jags/JAGS/4.x/Source/JAGS-4.3.2.zip"
	jags_cmd_zip := exec.Command("curl", "-L", jags_zip_url, "-o", "package.zip")
	err = jags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jags_bin_url := "https://downloads.sourceforge.net/project/mcmc-jags/JAGS/4.x/Source/JAGS-4.3.2.bin"
	jags_cmd_bin := exec.Command("curl", "-L", jags_bin_url, "-o", "binary.bin")
	err = jags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jags_src_url := "https://downloads.sourceforge.net/project/mcmc-jags/JAGS/4.x/Source/JAGS-4.3.2.src.tar.gz"
	jags_cmd_src := exec.Command("curl", "-L", jags_src_url, "-o", "source.tar.gz")
	err = jags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jags_cmd_direct := exec.Command("./binary")
	err = jags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
