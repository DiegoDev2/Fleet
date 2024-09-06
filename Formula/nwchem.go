package main

// Nwchem - High-performance computational chemistry tools
// Homepage: https://nwchemgit.github.io

import (
	"fmt"
	
	"os/exec"
)

func installNwchem() {
	// Método 1: Descargar y extraer .tar.gz
	nwchem_tar_url := "https://github.com/nwchemgit/nwchem/releases/download/v7.2.3-release/nwchem-7.2.3-release.revision-d690e065-src.2024-08-27.tar.xz"
	nwchem_cmd_tar := exec.Command("curl", "-L", nwchem_tar_url, "-o", "package.tar.gz")
	err := nwchem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nwchem_zip_url := "https://github.com/nwchemgit/nwchem/releases/download/v7.2.3-release/nwchem-7.2.3-release.revision-d690e065-src.2024-08-27.tar.xz"
	nwchem_cmd_zip := exec.Command("curl", "-L", nwchem_zip_url, "-o", "package.zip")
	err = nwchem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nwchem_bin_url := "https://github.com/nwchemgit/nwchem/releases/download/v7.2.3-release/nwchem-7.2.3-release.revision-d690e065-src.2024-08-27.tar.xz"
	nwchem_cmd_bin := exec.Command("curl", "-L", nwchem_bin_url, "-o", "binary.bin")
	err = nwchem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nwchem_src_url := "https://github.com/nwchemgit/nwchem/releases/download/v7.2.3-release/nwchem-7.2.3-release.revision-d690e065-src.2024-08-27.tar.xz"
	nwchem_cmd_src := exec.Command("curl", "-L", nwchem_src_url, "-o", "source.tar.gz")
	err = nwchem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nwchem_cmd_direct := exec.Command("./binary")
	err = nwchem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: hwloc")
	exec.Command("latte", "install", "hwloc").Run()
	fmt.Println("Instalando dependencia: libxc")
	exec.Command("latte", "install", "libxc").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: scalapack")
	exec.Command("latte", "install", "scalapack").Run()
}
