package main

// Dnstwist - Test domains for typo squatting, phishing and corporate espionage
// Homepage: https://github.com/elceef/dnstwist

import (
	"fmt"
	
	"os/exec"
)

func installDnstwist() {
	// Método 1: Descargar y extraer .tar.gz
	dnstwist_tar_url := "https://files.pythonhosted.org/packages/3f/df/9c62d9e40d374fd1311de3c761670771615101e0a0b31968b31289882db7/dnstwist-20240812.tar.gz"
	dnstwist_cmd_tar := exec.Command("curl", "-L", dnstwist_tar_url, "-o", "package.tar.gz")
	err := dnstwist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnstwist_zip_url := "https://files.pythonhosted.org/packages/3f/df/9c62d9e40d374fd1311de3c761670771615101e0a0b31968b31289882db7/dnstwist-20240812.zip"
	dnstwist_cmd_zip := exec.Command("curl", "-L", dnstwist_zip_url, "-o", "package.zip")
	err = dnstwist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnstwist_bin_url := "https://files.pythonhosted.org/packages/3f/df/9c62d9e40d374fd1311de3c761670771615101e0a0b31968b31289882db7/dnstwist-20240812.bin"
	dnstwist_cmd_bin := exec.Command("curl", "-L", dnstwist_bin_url, "-o", "binary.bin")
	err = dnstwist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnstwist_src_url := "https://files.pythonhosted.org/packages/3f/df/9c62d9e40d374fd1311de3c761670771615101e0a0b31968b31289882db7/dnstwist-20240812.src.tar.gz"
	dnstwist_cmd_src := exec.Command("curl", "-L", dnstwist_src_url, "-o", "source.tar.gz")
	err = dnstwist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnstwist_cmd_direct := exec.Command("./binary")
	err = dnstwist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libmaxminddb")
exec.Command("latte", "install", "libmaxminddb")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ssdeep")
exec.Command("latte", "install", "ssdeep")
	fmt.Println("Instalando dependencia: whois")
exec.Command("latte", "install", "whois")
}
