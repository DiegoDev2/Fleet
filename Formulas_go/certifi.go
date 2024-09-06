package main

// Certifi - Mozilla CA bundle for Python
// Homepage: https://github.com/certifi/python-certifi

import (
	"fmt"
	
	"os/exec"
)

func installCertifi() {
	// Método 1: Descargar y extraer .tar.gz
	certifi_tar_url := "https://files.pythonhosted.org/packages/b0/ee/9b19140fe824b367c04c5e1b369942dd754c4c5462d5674002f75c4dedc1/certifi-2024.8.30.tar.gz"
	certifi_cmd_tar := exec.Command("curl", "-L", certifi_tar_url, "-o", "package.tar.gz")
	err := certifi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	certifi_zip_url := "https://files.pythonhosted.org/packages/b0/ee/9b19140fe824b367c04c5e1b369942dd754c4c5462d5674002f75c4dedc1/certifi-2024.8.30.zip"
	certifi_cmd_zip := exec.Command("curl", "-L", certifi_zip_url, "-o", "package.zip")
	err = certifi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	certifi_bin_url := "https://files.pythonhosted.org/packages/b0/ee/9b19140fe824b367c04c5e1b369942dd754c4c5462d5674002f75c4dedc1/certifi-2024.8.30.bin"
	certifi_cmd_bin := exec.Command("curl", "-L", certifi_bin_url, "-o", "binary.bin")
	err = certifi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	certifi_src_url := "https://files.pythonhosted.org/packages/b0/ee/9b19140fe824b367c04c5e1b369942dd754c4c5462d5674002f75c4dedc1/certifi-2024.8.30.src.tar.gz"
	certifi_cmd_src := exec.Command("curl", "-L", certifi_src_url, "-o", "source.tar.gz")
	err = certifi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	certifi_cmd_direct := exec.Command("./binary")
	err = certifi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ca-certificates")
exec.Command("latte", "install", "ca-certificates")
}
