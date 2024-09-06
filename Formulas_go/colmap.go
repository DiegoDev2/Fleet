package main

// Colmap - Structure-from-Motion and Multi-View Stereo
// Homepage: https://colmap.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installColmap() {
	// Método 1: Descargar y extraer .tar.gz
	colmap_tar_url := "https://github.com/colmap/colmap/archive/refs/tags/3.10.tar.gz"
	colmap_cmd_tar := exec.Command("curl", "-L", colmap_tar_url, "-o", "package.tar.gz")
	err := colmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colmap_zip_url := "https://github.com/colmap/colmap/archive/refs/tags/3.10.zip"
	colmap_cmd_zip := exec.Command("curl", "-L", colmap_zip_url, "-o", "package.zip")
	err = colmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colmap_bin_url := "https://github.com/colmap/colmap/archive/refs/tags/3.10.bin"
	colmap_cmd_bin := exec.Command("curl", "-L", colmap_bin_url, "-o", "binary.bin")
	err = colmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colmap_src_url := "https://github.com/colmap/colmap/archive/refs/tags/3.10.src.tar.gz"
	colmap_cmd_src := exec.Command("curl", "-L", colmap_src_url, "-o", "source.tar.gz")
	err = colmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colmap_cmd_direct := exec.Command("./binary")
	err = colmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: ceres-solver")
exec.Command("latte", "install", "ceres-solver")
	fmt.Println("Instalando dependencia: cgal")
exec.Command("latte", "install", "cgal")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: flann")
exec.Command("latte", "install", "flann")
	fmt.Println("Instalando dependencia: freeimage")
exec.Command("latte", "install", "freeimage")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: metis")
exec.Command("latte", "install", "metis")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: suite-sparse")
exec.Command("latte", "install", "suite-sparse")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
