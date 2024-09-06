package main

// Lexicon - Manipulate DNS records on various DNS providers in a standardized way
// Homepage: https://github.com/AnalogJ/lexicon

import (
	"fmt"
	
	"os/exec"
)

func installLexicon() {
	// Método 1: Descargar y extraer .tar.gz
	lexicon_tar_url := "https://files.pythonhosted.org/packages/23/2a/b489cf65dd67d004847aea671729c2d8430fd9889393e6b77bb21e1a1029/dns_lexicon-3.18.0.tar.gz"
	lexicon_cmd_tar := exec.Command("curl", "-L", lexicon_tar_url, "-o", "package.tar.gz")
	err := lexicon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lexicon_zip_url := "https://files.pythonhosted.org/packages/23/2a/b489cf65dd67d004847aea671729c2d8430fd9889393e6b77bb21e1a1029/dns_lexicon-3.18.0.zip"
	lexicon_cmd_zip := exec.Command("curl", "-L", lexicon_zip_url, "-o", "package.zip")
	err = lexicon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lexicon_bin_url := "https://files.pythonhosted.org/packages/23/2a/b489cf65dd67d004847aea671729c2d8430fd9889393e6b77bb21e1a1029/dns_lexicon-3.18.0.bin"
	lexicon_cmd_bin := exec.Command("curl", "-L", lexicon_bin_url, "-o", "binary.bin")
	err = lexicon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lexicon_src_url := "https://files.pythonhosted.org/packages/23/2a/b489cf65dd67d004847aea671729c2d8430fd9889393e6b77bb21e1a1029/dns_lexicon-3.18.0.src.tar.gz"
	lexicon_cmd_src := exec.Command("curl", "-L", lexicon_src_url, "-o", "source.tar.gz")
	err = lexicon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lexicon_cmd_direct := exec.Command("./binary")
	err = lexicon_cmd_direct.Run()
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
