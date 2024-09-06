package main

// CryticCompile - Abstraction layer for smart contract build systems
// Homepage: https://github.com/crytic/crytic-compile

import (
	"fmt"
	
	"os/exec"
)

func installCryticCompile() {
	// Método 1: Descargar y extraer .tar.gz
	cryticcompile_tar_url := "https://files.pythonhosted.org/packages/54/f8/6833fb37702900711e5617e0594e2eeccbb0b716993e84b00ee186907e1c/crytic-compile-0.3.7.tar.gz"
	cryticcompile_cmd_tar := exec.Command("curl", "-L", cryticcompile_tar_url, "-o", "package.tar.gz")
	err := cryticcompile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryticcompile_zip_url := "https://files.pythonhosted.org/packages/54/f8/6833fb37702900711e5617e0594e2eeccbb0b716993e84b00ee186907e1c/crytic-compile-0.3.7.zip"
	cryticcompile_cmd_zip := exec.Command("curl", "-L", cryticcompile_zip_url, "-o", "package.zip")
	err = cryticcompile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryticcompile_bin_url := "https://files.pythonhosted.org/packages/54/f8/6833fb37702900711e5617e0594e2eeccbb0b716993e84b00ee186907e1c/crytic-compile-0.3.7.bin"
	cryticcompile_cmd_bin := exec.Command("curl", "-L", cryticcompile_bin_url, "-o", "binary.bin")
	err = cryticcompile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryticcompile_src_url := "https://files.pythonhosted.org/packages/54/f8/6833fb37702900711e5617e0594e2eeccbb0b716993e84b00ee186907e1c/crytic-compile-0.3.7.src.tar.gz"
	cryticcompile_cmd_src := exec.Command("curl", "-L", cryticcompile_src_url, "-o", "source.tar.gz")
	err = cryticcompile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryticcompile_cmd_direct := exec.Command("./binary")
	err = cryticcompile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
