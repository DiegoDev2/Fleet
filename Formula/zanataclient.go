package main

// ZanataClient - Zanata translation system command-line client
// Homepage: http://zanata.org/

import (
	"fmt"
	
	"os/exec"
)

func installZanataClient() {
	// Método 1: Descargar y extraer .tar.gz
	zanataclient_tar_url := "https://search.maven.org/remotecontent?filepath=org/zanata/zanata-cli/4.6.2/zanata-cli-4.6.2-dist.tar.gz"
	zanataclient_cmd_tar := exec.Command("curl", "-L", zanataclient_tar_url, "-o", "package.tar.gz")
	err := zanataclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zanataclient_zip_url := "https://search.maven.org/remotecontent?filepath=org/zanata/zanata-cli/4.6.2/zanata-cli-4.6.2-dist.zip"
	zanataclient_cmd_zip := exec.Command("curl", "-L", zanataclient_zip_url, "-o", "package.zip")
	err = zanataclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zanataclient_bin_url := "https://search.maven.org/remotecontent?filepath=org/zanata/zanata-cli/4.6.2/zanata-cli-4.6.2-dist.bin"
	zanataclient_cmd_bin := exec.Command("curl", "-L", zanataclient_bin_url, "-o", "binary.bin")
	err = zanataclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zanataclient_src_url := "https://search.maven.org/remotecontent?filepath=org/zanata/zanata-cli/4.6.2/zanata-cli-4.6.2-dist.src.tar.gz"
	zanataclient_cmd_src := exec.Command("curl", "-L", zanataclient_src_url, "-o", "source.tar.gz")
	err = zanataclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zanataclient_cmd_direct := exec.Command("./binary")
	err = zanataclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
	exec.Command("latte", "install", "openjdk@8").Run()
}
