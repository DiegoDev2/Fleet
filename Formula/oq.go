package main

// Oq - Performant, and portable jq wrapper to support formats other than JSON
// Homepage: https://blacksmoke16.github.io/oq

import (
	"fmt"
	
	"os/exec"
)

func installOq() {
	// Método 1: Descargar y extraer .tar.gz
	oq_tar_url := "https://github.com/Blacksmoke16/oq/archive/refs/tags/v1.3.5.tar.gz"
	oq_cmd_tar := exec.Command("curl", "-L", oq_tar_url, "-o", "package.tar.gz")
	err := oq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oq_zip_url := "https://github.com/Blacksmoke16/oq/archive/refs/tags/v1.3.5.zip"
	oq_cmd_zip := exec.Command("curl", "-L", oq_zip_url, "-o", "package.zip")
	err = oq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oq_bin_url := "https://github.com/Blacksmoke16/oq/archive/refs/tags/v1.3.5.bin"
	oq_cmd_bin := exec.Command("curl", "-L", oq_bin_url, "-o", "binary.bin")
	err = oq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oq_src_url := "https://github.com/Blacksmoke16/oq/archive/refs/tags/v1.3.5.src.tar.gz"
	oq_cmd_src := exec.Command("curl", "-L", oq_src_url, "-o", "source.tar.gz")
	err = oq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oq_cmd_direct := exec.Command("./binary")
	err = oq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: crystal")
	exec.Command("latte", "install", "crystal").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
