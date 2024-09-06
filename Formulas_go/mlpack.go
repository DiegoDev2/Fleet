package main

// Mlpack - Scalable C++ machine learning library
// Homepage: https://www.mlpack.org

import (
	"fmt"
	
	"os/exec"
)

func installMlpack() {
	// Método 1: Descargar y extraer .tar.gz
	mlpack_tar_url := "https://mlpack.org/files/mlpack-4.4.0.tar.gz"
	mlpack_cmd_tar := exec.Command("curl", "-L", mlpack_tar_url, "-o", "package.tar.gz")
	err := mlpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlpack_zip_url := "https://mlpack.org/files/mlpack-4.4.0.zip"
	mlpack_cmd_zip := exec.Command("curl", "-L", mlpack_zip_url, "-o", "package.zip")
	err = mlpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlpack_bin_url := "https://mlpack.org/files/mlpack-4.4.0.bin"
	mlpack_cmd_bin := exec.Command("curl", "-L", mlpack_bin_url, "-o", "binary.bin")
	err = mlpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlpack_src_url := "https://mlpack.org/files/mlpack-4.4.0.src.tar.gz"
	mlpack_cmd_src := exec.Command("curl", "-L", mlpack_src_url, "-o", "source.tar.gz")
	err = mlpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlpack_cmd_direct := exec.Command("./binary")
	err = mlpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: armadillo")
exec.Command("latte", "install", "armadillo")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cereal")
exec.Command("latte", "install", "cereal")
	fmt.Println("Instalando dependencia: ensmallen")
exec.Command("latte", "install", "ensmallen")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
}
