package main

// Alpscore - Applications and libraries for physics simulations
// Homepage: https://alpscore.org

import (
	"fmt"
	
	"os/exec"
)

func installAlpscore() {
	// Método 1: Descargar y extraer .tar.gz
	alpscore_tar_url := "https://github.com/ALPSCore/ALPSCore/archive/refs/tags/v2.3.1.tar.gz"
	alpscore_cmd_tar := exec.Command("curl", "-L", alpscore_tar_url, "-o", "package.tar.gz")
	err := alpscore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alpscore_zip_url := "https://github.com/ALPSCore/ALPSCore/archive/refs/tags/v2.3.1.zip"
	alpscore_cmd_zip := exec.Command("curl", "-L", alpscore_zip_url, "-o", "package.zip")
	err = alpscore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alpscore_bin_url := "https://github.com/ALPSCore/ALPSCore/archive/refs/tags/v2.3.1.bin"
	alpscore_cmd_bin := exec.Command("curl", "-L", alpscore_bin_url, "-o", "binary.bin")
	err = alpscore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alpscore_src_url := "https://github.com/ALPSCore/ALPSCore/archive/refs/tags/v2.3.1.src.tar.gz"
	alpscore_cmd_src := exec.Command("curl", "-L", alpscore_src_url, "-o", "source.tar.gz")
	err = alpscore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alpscore_cmd_direct := exec.Command("./binary")
	err = alpscore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
