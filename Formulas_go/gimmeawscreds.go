package main

// GimmeAwsCreds - CLI to retrieve AWS credentials from Okta
// Homepage: https://github.com/Nike-Inc/gimme-aws-creds

import (
	"fmt"
	
	"os/exec"
)

func installGimmeAwsCreds() {
	// Método 1: Descargar y extraer .tar.gz
	gimmeawscreds_tar_url := "https://files.pythonhosted.org/packages/63/73/9e508d37d4d301f6a3811fdc0b0a076696de87f82ad8a81ec28c3e6befb5/gimme_aws_creds-2.8.2.tar.gz"
	gimmeawscreds_cmd_tar := exec.Command("curl", "-L", gimmeawscreds_tar_url, "-o", "package.tar.gz")
	err := gimmeawscreds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gimmeawscreds_zip_url := "https://files.pythonhosted.org/packages/63/73/9e508d37d4d301f6a3811fdc0b0a076696de87f82ad8a81ec28c3e6befb5/gimme_aws_creds-2.8.2.zip"
	gimmeawscreds_cmd_zip := exec.Command("curl", "-L", gimmeawscreds_zip_url, "-o", "package.zip")
	err = gimmeawscreds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gimmeawscreds_bin_url := "https://files.pythonhosted.org/packages/63/73/9e508d37d4d301f6a3811fdc0b0a076696de87f82ad8a81ec28c3e6befb5/gimme_aws_creds-2.8.2.bin"
	gimmeawscreds_cmd_bin := exec.Command("curl", "-L", gimmeawscreds_bin_url, "-o", "binary.bin")
	err = gimmeawscreds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gimmeawscreds_src_url := "https://files.pythonhosted.org/packages/63/73/9e508d37d4d301f6a3811fdc0b0a076696de87f82ad8a81ec28c3e6befb5/gimme_aws_creds-2.8.2.src.tar.gz"
	gimmeawscreds_cmd_src := exec.Command("curl", "-L", gimmeawscreds_src_url, "-o", "source.tar.gz")
	err = gimmeawscreds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gimmeawscreds_cmd_direct := exec.Command("./binary")
	err = gimmeawscreds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
