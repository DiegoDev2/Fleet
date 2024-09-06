package main

// Tfsec - Static analysis security scanner for your terraform code
// Homepage: https://aquasecurity.github.io/tfsec/latest/

import (
	"fmt"
	
	"os/exec"
)

func installTfsec() {
	// Método 1: Descargar y extraer .tar.gz
	tfsec_tar_url := "https://github.com/aquasecurity/tfsec/archive/refs/tags/v1.28.10.tar.gz"
	tfsec_cmd_tar := exec.Command("curl", "-L", tfsec_tar_url, "-o", "package.tar.gz")
	err := tfsec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfsec_zip_url := "https://github.com/aquasecurity/tfsec/archive/refs/tags/v1.28.10.zip"
	tfsec_cmd_zip := exec.Command("curl", "-L", tfsec_zip_url, "-o", "package.zip")
	err = tfsec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfsec_bin_url := "https://github.com/aquasecurity/tfsec/archive/refs/tags/v1.28.10.bin"
	tfsec_cmd_bin := exec.Command("curl", "-L", tfsec_bin_url, "-o", "binary.bin")
	err = tfsec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfsec_src_url := "https://github.com/aquasecurity/tfsec/archive/refs/tags/v1.28.10.src.tar.gz"
	tfsec_cmd_src := exec.Command("curl", "-L", tfsec_src_url, "-o", "source.tar.gz")
	err = tfsec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfsec_cmd_direct := exec.Command("./binary")
	err = tfsec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
