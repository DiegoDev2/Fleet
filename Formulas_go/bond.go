package main

// Bond - Cross-platform framework for working with schematized data
// Homepage: https://github.com/microsoft/bond

import (
	"fmt"
	
	"os/exec"
)

func installBond() {
	// Método 1: Descargar y extraer .tar.gz
	bond_tar_url := "https://github.com/microsoft/bond/archive/refs/tags/10.0.0.tar.gz"
	bond_cmd_tar := exec.Command("curl", "-L", bond_tar_url, "-o", "package.tar.gz")
	err := bond_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bond_zip_url := "https://github.com/microsoft/bond/archive/refs/tags/10.0.0.zip"
	bond_cmd_zip := exec.Command("curl", "-L", bond_zip_url, "-o", "package.zip")
	err = bond_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bond_bin_url := "https://github.com/microsoft/bond/archive/refs/tags/10.0.0.bin"
	bond_cmd_bin := exec.Command("curl", "-L", bond_bin_url, "-o", "binary.bin")
	err = bond_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bond_src_url := "https://github.com/microsoft/bond/archive/refs/tags/10.0.0.src.tar.gz"
	bond_cmd_src := exec.Command("curl", "-L", bond_src_url, "-o", "source.tar.gz")
	err = bond_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bond_cmd_direct := exec.Command("./binary")
	err = bond_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ghc@8.6")
exec.Command("latte", "install", "ghc@8.6")
	fmt.Println("Instalando dependencia: haskell-stack")
exec.Command("latte", "install", "haskell-stack")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: rapidjson")
exec.Command("latte", "install", "rapidjson")
}
