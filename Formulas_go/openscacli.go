package main

// OpenscaCli - OpenSCA is a supply-chain security tool for security researchers and developers
// Homepage: https://opensca.xmirror.cn

import (
	"fmt"
	
	"os/exec"
)

func installOpenscaCli() {
	// Método 1: Descargar y extraer .tar.gz
	openscacli_tar_url := "https://github.com/XmirrorSecurity/OpenSCA-cli/archive/refs/tags/v3.0.5.tar.gz"
	openscacli_cmd_tar := exec.Command("curl", "-L", openscacli_tar_url, "-o", "package.tar.gz")
	err := openscacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openscacli_zip_url := "https://github.com/XmirrorSecurity/OpenSCA-cli/archive/refs/tags/v3.0.5.zip"
	openscacli_cmd_zip := exec.Command("curl", "-L", openscacli_zip_url, "-o", "package.zip")
	err = openscacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openscacli_bin_url := "https://github.com/XmirrorSecurity/OpenSCA-cli/archive/refs/tags/v3.0.5.bin"
	openscacli_cmd_bin := exec.Command("curl", "-L", openscacli_bin_url, "-o", "binary.bin")
	err = openscacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openscacli_src_url := "https://github.com/XmirrorSecurity/OpenSCA-cli/archive/refs/tags/v3.0.5.src.tar.gz"
	openscacli_cmd_src := exec.Command("curl", "-L", openscacli_src_url, "-o", "source.tar.gz")
	err = openscacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openscacli_cmd_direct := exec.Command("./binary")
	err = openscacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
