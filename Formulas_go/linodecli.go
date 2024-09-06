package main

// LinodeCli - CLI for the Linode API
// Homepage: https://github.com/linode/linode-cli

import (
	"fmt"
	
	"os/exec"
)

func installLinodeCli() {
	// Método 1: Descargar y extraer .tar.gz
	linodecli_tar_url := "https://files.pythonhosted.org/packages/e9/09/6fb0d3ccafcd3f784a9f8cb0dab43ff284dd761b1a6f7609dfcb2213f713/linode_cli-5.51.0.tar.gz"
	linodecli_cmd_tar := exec.Command("curl", "-L", linodecli_tar_url, "-o", "package.tar.gz")
	err := linodecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linodecli_zip_url := "https://files.pythonhosted.org/packages/e9/09/6fb0d3ccafcd3f784a9f8cb0dab43ff284dd761b1a6f7609dfcb2213f713/linode_cli-5.51.0.zip"
	linodecli_cmd_zip := exec.Command("curl", "-L", linodecli_zip_url, "-o", "package.zip")
	err = linodecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linodecli_bin_url := "https://files.pythonhosted.org/packages/e9/09/6fb0d3ccafcd3f784a9f8cb0dab43ff284dd761b1a6f7609dfcb2213f713/linode_cli-5.51.0.bin"
	linodecli_cmd_bin := exec.Command("curl", "-L", linodecli_bin_url, "-o", "binary.bin")
	err = linodecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linodecli_src_url := "https://files.pythonhosted.org/packages/e9/09/6fb0d3ccafcd3f784a9f8cb0dab43ff284dd761b1a6f7609dfcb2213f713/linode_cli-5.51.0.src.tar.gz"
	linodecli_cmd_src := exec.Command("curl", "-L", linodecli_src_url, "-o", "source.tar.gz")
	err = linodecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linodecli_cmd_direct := exec.Command("./binary")
	err = linodecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
