package main

// ApacheOpennlp - Machine learning toolkit for processing natural language text
// Homepage: https://opennlp.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installApacheOpennlp() {
	// Método 1: Descargar y extraer .tar.gz
	apacheopennlp_tar_url := "https://www.apache.org/dyn/closer.lua?path=opennlp/opennlp-2.4.0/apache-opennlp-2.4.0-bin.tar.gz"
	apacheopennlp_cmd_tar := exec.Command("curl", "-L", apacheopennlp_tar_url, "-o", "package.tar.gz")
	err := apacheopennlp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apacheopennlp_zip_url := "https://www.apache.org/dyn/closer.lua?path=opennlp/opennlp-2.4.0/apache-opennlp-2.4.0-bin.zip"
	apacheopennlp_cmd_zip := exec.Command("curl", "-L", apacheopennlp_zip_url, "-o", "package.zip")
	err = apacheopennlp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apacheopennlp_bin_url := "https://www.apache.org/dyn/closer.lua?path=opennlp/opennlp-2.4.0/apache-opennlp-2.4.0-bin.bin"
	apacheopennlp_cmd_bin := exec.Command("curl", "-L", apacheopennlp_bin_url, "-o", "binary.bin")
	err = apacheopennlp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apacheopennlp_src_url := "https://www.apache.org/dyn/closer.lua?path=opennlp/opennlp-2.4.0/apache-opennlp-2.4.0-bin.src.tar.gz"
	apacheopennlp_cmd_src := exec.Command("curl", "-L", apacheopennlp_src_url, "-o", "source.tar.gz")
	err = apacheopennlp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apacheopennlp_cmd_direct := exec.Command("./binary")
	err = apacheopennlp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
