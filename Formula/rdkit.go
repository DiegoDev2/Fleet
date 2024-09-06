package main

// Rdkit - Open-source chemoinformatics library
// Homepage: https://rdkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installRdkit() {
	// Método 1: Descargar y extraer .tar.gz
	rdkit_tar_url := "https://github.com/rdkit/rdkit/archive/refs/tags/Release_2024_03_6.tar.gz"
	rdkit_cmd_tar := exec.Command("curl", "-L", rdkit_tar_url, "-o", "package.tar.gz")
	err := rdkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rdkit_zip_url := "https://github.com/rdkit/rdkit/archive/refs/tags/Release_2024_03_6.zip"
	rdkit_cmd_zip := exec.Command("curl", "-L", rdkit_zip_url, "-o", "package.zip")
	err = rdkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rdkit_bin_url := "https://github.com/rdkit/rdkit/archive/refs/tags/Release_2024_03_6.bin"
	rdkit_cmd_bin := exec.Command("curl", "-L", rdkit_bin_url, "-o", "binary.bin")
	err = rdkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rdkit_src_url := "https://github.com/rdkit/rdkit/archive/refs/tags/Release_2024_03_6.src.tar.gz"
	rdkit_cmd_src := exec.Command("curl", "-L", rdkit_src_url, "-o", "source.tar.gz")
	err = rdkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rdkit_cmd_direct := exec.Command("./binary")
	err = rdkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: catch2")
	exec.Command("latte", "install", "catch2").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
	fmt.Println("Instalando dependencia: py3cairo")
	exec.Command("latte", "install", "py3cairo").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
