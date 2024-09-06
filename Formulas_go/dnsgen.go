package main

// Dnsgen - Generates DNS names from existing domain names
// Homepage: https://github.com/AlephNullSK/dnsgen

import (
	"fmt"
	
	"os/exec"
)

func installDnsgen() {
	// Método 1: Descargar y extraer .tar.gz
	dnsgen_tar_url := "https://files.pythonhosted.org/packages/5f/e1/1c7d86f51da5b93f3f99ac99e3ad051ed82234147ddd869f77a3959e6abc/dnsgen-1.0.4.tar.gz"
	dnsgen_cmd_tar := exec.Command("curl", "-L", dnsgen_tar_url, "-o", "package.tar.gz")
	err := dnsgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsgen_zip_url := "https://files.pythonhosted.org/packages/5f/e1/1c7d86f51da5b93f3f99ac99e3ad051ed82234147ddd869f77a3959e6abc/dnsgen-1.0.4.zip"
	dnsgen_cmd_zip := exec.Command("curl", "-L", dnsgen_zip_url, "-o", "package.zip")
	err = dnsgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsgen_bin_url := "https://files.pythonhosted.org/packages/5f/e1/1c7d86f51da5b93f3f99ac99e3ad051ed82234147ddd869f77a3959e6abc/dnsgen-1.0.4.bin"
	dnsgen_cmd_bin := exec.Command("curl", "-L", dnsgen_bin_url, "-o", "binary.bin")
	err = dnsgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsgen_src_url := "https://files.pythonhosted.org/packages/5f/e1/1c7d86f51da5b93f3f99ac99e3ad051ed82234147ddd869f77a3959e6abc/dnsgen-1.0.4.src.tar.gz"
	dnsgen_cmd_src := exec.Command("curl", "-L", dnsgen_src_url, "-o", "source.tar.gz")
	err = dnsgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsgen_cmd_direct := exec.Command("./binary")
	err = dnsgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
