package main

// Lammps - Molecular Dynamics Simulator
// Homepage: https://docs.lammps.org/

import (
	"fmt"
	
	"os/exec"
)

func installLammps() {
	// Método 1: Descargar y extraer .tar.gz
	lammps_tar_url := "https://github.com/lammps/lammps/archive/refs/tags/stable_29Aug2024.tar.gz"
	lammps_cmd_tar := exec.Command("curl", "-L", lammps_tar_url, "-o", "package.tar.gz")
	err := lammps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lammps_zip_url := "https://github.com/lammps/lammps/archive/refs/tags/stable_29Aug2024.zip"
	lammps_cmd_zip := exec.Command("curl", "-L", lammps_zip_url, "-o", "package.zip")
	err = lammps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lammps_bin_url := "https://github.com/lammps/lammps/archive/refs/tags/stable_29Aug2024.bin"
	lammps_cmd_bin := exec.Command("curl", "-L", lammps_bin_url, "-o", "binary.bin")
	err = lammps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lammps_src_url := "https://github.com/lammps/lammps/archive/refs/tags/stable_29Aug2024.src.tar.gz"
	lammps_cmd_src := exec.Command("curl", "-L", lammps_src_url, "-o", "source.tar.gz")
	err = lammps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lammps_cmd_direct := exec.Command("./binary")
	err = lammps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: gsl")
	exec.Command("latte", "install", "gsl").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: kim-api")
	exec.Command("latte", "install", "kim-api").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
