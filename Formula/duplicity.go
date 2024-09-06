package main

// Duplicity - Bandwidth-efficient encrypted backup
// Homepage: https://gitlab.com/duplicity/duplicity

import (
	"fmt"
	
	"os/exec"
)

func installDuplicity() {
	// Método 1: Descargar y extraer .tar.gz
	duplicity_tar_url := "https://files.pythonhosted.org/packages/c5/a6/99dc9081acdfbbb1d3c8cb557f468e5ac9bbf7b976472cc9e463d427ad30/duplicity-3.0.2.tar.gz"
	duplicity_cmd_tar := exec.Command("curl", "-L", duplicity_tar_url, "-o", "package.tar.gz")
	err := duplicity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duplicity_zip_url := "https://files.pythonhosted.org/packages/c5/a6/99dc9081acdfbbb1d3c8cb557f468e5ac9bbf7b976472cc9e463d427ad30/duplicity-3.0.2.zip"
	duplicity_cmd_zip := exec.Command("curl", "-L", duplicity_zip_url, "-o", "package.zip")
	err = duplicity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duplicity_bin_url := "https://files.pythonhosted.org/packages/c5/a6/99dc9081acdfbbb1d3c8cb557f468e5ac9bbf7b976472cc9e463d427ad30/duplicity-3.0.2.bin"
	duplicity_cmd_bin := exec.Command("curl", "-L", duplicity_bin_url, "-o", "binary.bin")
	err = duplicity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duplicity_src_url := "https://files.pythonhosted.org/packages/c5/a6/99dc9081acdfbbb1d3c8cb557f468e5ac9bbf7b976472cc9e463d427ad30/duplicity-3.0.2.src.tar.gz"
	duplicity_cmd_src := exec.Command("curl", "-L", duplicity_src_url, "-o", "source.tar.gz")
	err = duplicity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duplicity_cmd_direct := exec.Command("./binary")
	err = duplicity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: librsync")
	exec.Command("latte", "install", "librsync").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
