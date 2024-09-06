package main

// OciCli - Oracle Cloud Infrastructure CLI
// Homepage: https://docs.cloud.oracle.com/iaas/Content/API/Concepts/cliconcepts.htm

import (
	"fmt"
	
	"os/exec"
)

func installOciCli() {
	// Método 1: Descargar y extraer .tar.gz
	ocicli_tar_url := "https://files.pythonhosted.org/packages/d5/93/0e0ffe2682f4625aa05f60882a481aac2908153f0d62a798a14364543b89/oci-cli-3.47.0.tar.gz"
	ocicli_cmd_tar := exec.Command("curl", "-L", ocicli_tar_url, "-o", "package.tar.gz")
	err := ocicli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocicli_zip_url := "https://files.pythonhosted.org/packages/d5/93/0e0ffe2682f4625aa05f60882a481aac2908153f0d62a798a14364543b89/oci-cli-3.47.0.zip"
	ocicli_cmd_zip := exec.Command("curl", "-L", ocicli_zip_url, "-o", "package.zip")
	err = ocicli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocicli_bin_url := "https://files.pythonhosted.org/packages/d5/93/0e0ffe2682f4625aa05f60882a481aac2908153f0d62a798a14364543b89/oci-cli-3.47.0.bin"
	ocicli_cmd_bin := exec.Command("curl", "-L", ocicli_bin_url, "-o", "binary.bin")
	err = ocicli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocicli_src_url := "https://files.pythonhosted.org/packages/d5/93/0e0ffe2682f4625aa05f60882a481aac2908153f0d62a798a14364543b89/oci-cli-3.47.0.src.tar.gz"
	ocicli_cmd_src := exec.Command("curl", "-L", ocicli_src_url, "-o", "source.tar.gz")
	err = ocicli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocicli_cmd_direct := exec.Command("./binary")
	err = ocicli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
