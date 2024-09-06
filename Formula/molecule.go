package main

// Molecule - Automated testing for Ansible roles
// Homepage: https://molecule.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installMolecule() {
	// Método 1: Descargar y extraer .tar.gz
	molecule_tar_url := "https://files.pythonhosted.org/packages/a2/01/ebbe0ace407b82439abde26c987d477b3b116776ddf554fec1a1b3542597/molecule-24.8.0.tar.gz"
	molecule_cmd_tar := exec.Command("curl", "-L", molecule_tar_url, "-o", "package.tar.gz")
	err := molecule_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	molecule_zip_url := "https://files.pythonhosted.org/packages/a2/01/ebbe0ace407b82439abde26c987d477b3b116776ddf554fec1a1b3542597/molecule-24.8.0.zip"
	molecule_cmd_zip := exec.Command("curl", "-L", molecule_zip_url, "-o", "package.zip")
	err = molecule_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	molecule_bin_url := "https://files.pythonhosted.org/packages/a2/01/ebbe0ace407b82439abde26c987d477b3b116776ddf554fec1a1b3542597/molecule-24.8.0.bin"
	molecule_cmd_bin := exec.Command("curl", "-L", molecule_bin_url, "-o", "binary.bin")
	err = molecule_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	molecule_src_url := "https://files.pythonhosted.org/packages/a2/01/ebbe0ace407b82439abde26c987d477b3b116776ddf554fec1a1b3542597/molecule-24.8.0.src.tar.gz"
	molecule_cmd_src := exec.Command("curl", "-L", molecule_src_url, "-o", "source.tar.gz")
	err = molecule_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	molecule_cmd_direct := exec.Command("./binary")
	err = molecule_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: ansible")
	exec.Command("latte", "install", "ansible").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
