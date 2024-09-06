package main

// Prowler - Tool for cloud security assessments, audits, incident response, and more
// Homepage: https://prowler.com/

import (
	"fmt"
	
	"os/exec"
)

func installProwler() {
	// Método 1: Descargar y extraer .tar.gz
	prowler_tar_url := "https://files.pythonhosted.org/packages/c3/42/acb1219feaaa621b2e374405e49c62144a2e9ed00aa9137f36ae4cd34d1b/prowler-4.3.5.tar.gz"
	prowler_cmd_tar := exec.Command("curl", "-L", prowler_tar_url, "-o", "package.tar.gz")
	err := prowler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prowler_zip_url := "https://files.pythonhosted.org/packages/c3/42/acb1219feaaa621b2e374405e49c62144a2e9ed00aa9137f36ae4cd34d1b/prowler-4.3.5.zip"
	prowler_cmd_zip := exec.Command("curl", "-L", prowler_zip_url, "-o", "package.zip")
	err = prowler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prowler_bin_url := "https://files.pythonhosted.org/packages/c3/42/acb1219feaaa621b2e374405e49c62144a2e9ed00aa9137f36ae4cd34d1b/prowler-4.3.5.bin"
	prowler_cmd_bin := exec.Command("curl", "-L", prowler_bin_url, "-o", "binary.bin")
	err = prowler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prowler_src_url := "https://files.pythonhosted.org/packages/c3/42/acb1219feaaa621b2e374405e49c62144a2e9ed00aa9137f36ae4cd34d1b/prowler-4.3.5.src.tar.gz"
	prowler_cmd_src := exec.Command("curl", "-L", prowler_src_url, "-o", "source.tar.gz")
	err = prowler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prowler_cmd_direct := exec.Command("./binary")
	err = prowler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
