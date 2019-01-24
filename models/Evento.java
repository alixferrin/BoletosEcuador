/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package models;



import java.util.Date;
/**
 *
 * @author johancanales
 */
public class Evento {
    private int idEvento;
    private boolean ticketsDisponibles;
    private String nombre;
    private String artista;
    private Date fecha;
    private Venue venueAsignado;

    public Evento(int idEvento, boolean ticketsDisponibles, String nombre, String artista, Date fecha) {
        this.idEvento = idEvento;
        this.ticketsDisponibles = ticketsDisponibles;
        this.nombre = nombre;
        this.artista = artista;
        this.fecha = fecha;
    }
    
    
    public int getIdEvento() {
        return idEvento;
    }

    public void setIdEvento(int idEvento) {
        this.idEvento = idEvento;
    }

    public boolean isTicketsDisponibles() {
        return ticketsDisponibles;
    }

    public void setTicketsDisponibles(boolean ticketsDisponibles) {
        this.ticketsDisponibles = ticketsDisponibles;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getArtista() {
        return artista;
    }

    public void setArtista(String artista) {
        this.artista = artista;
    }

    public Date getFecha() {
        return fecha;
    }

    public void setFecha(Date fecha) {
        this.fecha = fecha;
    }
    
    



}
