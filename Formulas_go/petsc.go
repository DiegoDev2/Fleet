package main

// Petsc - Portable, Extensible Toolkit for Scientific Computation (real)
// Homepage: https://petsc.org/

import (
	"fmt"
	
	"os/exec"
)

func installPetsc() {
	// Método 1: Descargar y extraer .tar.gz
	petsc_tar_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.tar.gz"
	petsc_cmd_tar := exec.Command("curl", "-L", petsc_tar_url, "-o", "package.tar.gz")
	err := petsc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	petsc_zip_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.zip"
	petsc_cmd_zip := exec.Command("curl", "-L", petsc_zip_url, "-o", "package.zip")
	err = petsc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	petsc_bin_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.bin"
	petsc_cmd_bin := exec.Command("curl", "-L", petsc_bin_url, "-o", "binary.bin")
	err = petsc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	petsc_src_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.src.tar.gz"
	petsc_cmd_src := exec.Command("curl", "-L", petsc_src_url, "-o", "source.tar.gz")
	err = petsc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	petsc_cmd_direct := exec.Command("./binary")
	err = petsc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hdf5-mpi")
exec.Command("latte", "install", "hdf5-mpi")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: metis")
exec.Command("latte", "install", "metis")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: scalapack")
exec.Command("latte", "install", "scalapack")
	fmt.Println("Instalando dependencia: suite-sparse")
exec.Command("latte", "install", "suite-sparse")
}
