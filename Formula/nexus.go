package main

// Nexus - Repository manager for binary software components
// Homepage: https://www.sonatype.org/

import (
	"fmt"
	
	"os/exec"
)

func installNexus() {
	// Método 1: Descargar y extraer .tar.gz
	nexus_tar_url := "https://github.com/sonatype/nexus-public/archive/refs/tags/release-3.38.1-01.tar.gz"
	nexus_cmd_tar := exec.Command("curl", "-L", nexus_tar_url, "-o", "package.tar.gz")
	err := nexus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nexus_zip_url := "https://github.com/sonatype/nexus-public/archive/refs/tags/release-3.38.1-01.zip"
	nexus_cmd_zip := exec.Command("curl", "-L", nexus_zip_url, "-o", "package.zip")
	err = nexus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nexus_bin_url := "https://github.com/sonatype/nexus-public/archive/refs/tags/release-3.38.1-01.bin"
	nexus_cmd_bin := exec.Command("curl", "-L", nexus_bin_url, "-o", "binary.bin")
	err = nexus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nexus_src_url := "https://github.com/sonatype/nexus-public/archive/refs/tags/release-3.38.1-01.src.tar.gz"
	nexus_cmd_src := exec.Command("curl", "-L", nexus_src_url, "-o", "source.tar.gz")
	err = nexus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nexus_cmd_direct := exec.Command("./binary")
	err = nexus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk@8")
	exec.Command("latte", "install", "openjdk@8").Run()
}
