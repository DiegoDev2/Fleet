package main

// Amass - In-depth attack surface mapping and asset discovery
// Homepage: https://owasp.org/www-project-amass/

import (
	"fmt"
	
	"os/exec"
)

func installAmass() {
	// Método 1: Descargar y extraer .tar.gz
	amass_tar_url := "https://github.com/owasp-amass/amass/archive/refs/tags/v4.2.0.tar.gz"
	amass_cmd_tar := exec.Command("curl", "-L", amass_tar_url, "-o", "package.tar.gz")
	err := amass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amass_zip_url := "https://github.com/owasp-amass/amass/archive/refs/tags/v4.2.0.zip"
	amass_cmd_zip := exec.Command("curl", "-L", amass_zip_url, "-o", "package.zip")
	err = amass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amass_bin_url := "https://github.com/owasp-amass/amass/archive/refs/tags/v4.2.0.bin"
	amass_cmd_bin := exec.Command("curl", "-L", amass_bin_url, "-o", "binary.bin")
	err = amass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amass_src_url := "https://github.com/owasp-amass/amass/archive/refs/tags/v4.2.0.src.tar.gz"
	amass_cmd_src := exec.Command("curl", "-L", amass_src_url, "-o", "source.tar.gz")
	err = amass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amass_cmd_direct := exec.Command("./binary")
	err = amass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
