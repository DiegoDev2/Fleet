package main

// Gotestwaf - Tool for API and OWASP attack simulation
// Homepage: https://lab.wallarm.com/test-your-waf-before-hackers/

import (
	"fmt"
	
	"os/exec"
)

func installGotestwaf() {
	// Método 1: Descargar y extraer .tar.gz
	gotestwaf_tar_url := "https://github.com/wallarm/gotestwaf/archive/refs/tags/v0.5.5.tar.gz"
	gotestwaf_cmd_tar := exec.Command("curl", "-L", gotestwaf_tar_url, "-o", "package.tar.gz")
	err := gotestwaf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotestwaf_zip_url := "https://github.com/wallarm/gotestwaf/archive/refs/tags/v0.5.5.zip"
	gotestwaf_cmd_zip := exec.Command("curl", "-L", gotestwaf_zip_url, "-o", "package.zip")
	err = gotestwaf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotestwaf_bin_url := "https://github.com/wallarm/gotestwaf/archive/refs/tags/v0.5.5.bin"
	gotestwaf_cmd_bin := exec.Command("curl", "-L", gotestwaf_bin_url, "-o", "binary.bin")
	err = gotestwaf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotestwaf_src_url := "https://github.com/wallarm/gotestwaf/archive/refs/tags/v0.5.5.src.tar.gz"
	gotestwaf_cmd_src := exec.Command("curl", "-L", gotestwaf_src_url, "-o", "source.tar.gz")
	err = gotestwaf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotestwaf_cmd_direct := exec.Command("./binary")
	err = gotestwaf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
