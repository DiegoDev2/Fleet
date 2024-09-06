package main

// Gmsh - 3D finite element grid generator with CAD engine
// Homepage: https://gmsh.info/

import (
	"fmt"
	
	"os/exec"
)

func installGmsh() {
	// Método 1: Descargar y extraer .tar.gz
	gmsh_tar_url := "https://gmsh.info/src/gmsh-4.13.1-source.tgz"
	gmsh_cmd_tar := exec.Command("curl", "-L", gmsh_tar_url, "-o", "package.tar.gz")
	err := gmsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmsh_zip_url := "https://gmsh.info/src/gmsh-4.13.1-source.tgz"
	gmsh_cmd_zip := exec.Command("curl", "-L", gmsh_zip_url, "-o", "package.zip")
	err = gmsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmsh_bin_url := "https://gmsh.info/src/gmsh-4.13.1-source.tgz"
	gmsh_cmd_bin := exec.Command("curl", "-L", gmsh_bin_url, "-o", "binary.bin")
	err = gmsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmsh_src_url := "https://gmsh.info/src/gmsh-4.13.1-source.tgz"
	gmsh_cmd_src := exec.Command("curl", "-L", gmsh_src_url, "-o", "source.tar.gz")
	err = gmsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmsh_cmd_direct := exec.Command("./binary")
	err = gmsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fltk")
	exec.Command("latte", "install", "fltk").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: opencascade")
	exec.Command("latte", "install", "opencascade").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
