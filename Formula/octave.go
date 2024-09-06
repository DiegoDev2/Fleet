package main

// Octave - High-level interpreted language for numerical computing
// Homepage: https://octave.org/index.html

import (
	"fmt"
	
	"os/exec"
)

func installOctave() {
	// Método 1: Descargar y extraer .tar.gz
	octave_tar_url := "https://ftp.gnu.org/gnu/octave/octave-9.2.0.tar.xz"
	octave_cmd_tar := exec.Command("curl", "-L", octave_tar_url, "-o", "package.tar.gz")
	err := octave_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	octave_zip_url := "https://ftp.gnu.org/gnu/octave/octave-9.2.0.tar.xz"
	octave_cmd_zip := exec.Command("curl", "-L", octave_zip_url, "-o", "package.zip")
	err = octave_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	octave_bin_url := "https://ftp.gnu.org/gnu/octave/octave-9.2.0.tar.xz"
	octave_cmd_bin := exec.Command("curl", "-L", octave_bin_url, "-o", "binary.bin")
	err = octave_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	octave_src_url := "https://ftp.gnu.org/gnu/octave/octave-9.2.0.tar.xz"
	octave_cmd_src := exec.Command("curl", "-L", octave_src_url, "-o", "source.tar.gz")
	err = octave_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	octave_cmd_direct := exec.Command("./binary")
	err = octave_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: icoutils")
	exec.Command("latte", "install", "icoutils").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: arpack")
	exec.Command("latte", "install", "arpack").Run()
	fmt.Println("Instalando dependencia: epstool")
	exec.Command("latte", "install", "epstool").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: fig2dev")
	exec.Command("latte", "install", "fig2dev").Run()
	fmt.Println("Instalando dependencia: fltk")
	exec.Command("latte", "install", "fltk").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: gl2ps")
	exec.Command("latte", "install", "gl2ps").Run()
	fmt.Println("Instalando dependencia: glpk")
	exec.Command("latte", "install", "glpk").Run()
	fmt.Println("Instalando dependencia: graphicsmagick")
	exec.Command("latte", "install", "graphicsmagick").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
	fmt.Println("Instalando dependencia: pstoedit")
	exec.Command("latte", "install", "pstoedit").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: qrupdate")
	exec.Command("latte", "install", "qrupdate").Run()
	fmt.Println("Instalando dependencia: qscintilla2")
	exec.Command("latte", "install", "qscintilla2").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: rapidjson")
	exec.Command("latte", "install", "rapidjson").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
	fmt.Println("Instalando dependencia: sundials")
	exec.Command("latte", "install", "sundials").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
